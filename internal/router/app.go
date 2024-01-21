package router

import (
	"dqix/internal/types"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type app struct {
	dataPath   string
	data       data
	cssVersion string
}

type data struct {
	inventoryMap types.InventoryMap
	monsterMap   types.MonsterMap
	dataMap      map[string]types.DataKey
}

func (d data) GetQuickThing(id string) (thing types.Thing) {
	dataKey, ok := d.dataMap[id]
	if !ok {
		return nil
	}

	return dataKey
}

func (d data) GetThing(id string) (thing types.Thing) {
	dataKey, ok := d.dataMap[id]
	if !ok {
		return nil
	}

	switch dataKey.Structure {
	case "inventory":
		return d.inventoryMap.GetInventoryFromDataKey(dataKey)
	case "monsters":
		return d.monsterMap.GetMonsterFromDataKey(dataKey)
	default:
		return nil
	}
}

func (d data) GetInventory(id string) (inventory types.Inventory) {
	i := d.GetThing(id)

	inventory, _ = i.(types.Inventory)
	return
}

func (d data) GetMonster(id string) (monster types.Monster) {
	m := d.GetThing(id)

	monster, _ = m.(types.Monster)
	return
}

type RouterOption interface {
	apply(*app)
}

type dataPathOption struct {
	path string
}

func WithData(path string) RouterOption {
	return &dataPathOption{path: path}
}

func (o dataPathOption) apply(a *app) {
	a.data.inventoryMap = make(types.InventoryMap)
	a.data.monsterMap = make(types.MonsterMap)
	a.data.dataMap = make(map[string]types.DataKey)
	a.dataPath = o.path
}

func (a *app) loadData() error {
	if a.dataPath == "" {
		return nil
	}

	files, err := os.ReadDir(a.dataPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			pathParts := strings.Split(file.Name(), "/")
			switch pathParts[len(pathParts)-1] {
			case "inventory":
				if err := a.loadInventory(filepath.Join(a.dataPath, file.Name())); err != nil {
					return err
				}
			case "monsters":
				if err := a.loadMonsters(filepath.Join(a.dataPath, file.Name())); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (a *app) loadInventory(basePath string) error {
	return filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var classMap map[string]types.Inventory
		if err := json.Unmarshal(raw, &classMap); err != nil {
			return err
		}

		for _, inventory := range classMap {
			a.data.inventoryMap.AddInventory(inventory)
			a.data.dataMap[inventory.GetID()] = inventory.ToDataKey()
		}

		return nil
	})
}

func (a *app) loadMonsters(basePath string) error {
	return filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var monsterMap map[string]types.Monster
		if err := json.Unmarshal(raw, &monsterMap); err != nil {
			return err
		}

		for _, monster := range monsterMap {
			a.data.monsterMap.AddMonster(monster)
			a.data.dataMap[monster.GetID()] = monster.ToDataKey()
		}

		return nil
	})
}

type cssVersionOption struct {
	cssVersion string
}

func WithCSSVersion(cssVersion string) RouterOption {
	return &cssVersionOption{cssVersion: cssVersion}
}

func (o cssVersionOption) apply(a *app) {
	a.cssVersion = o.cssVersion
}
