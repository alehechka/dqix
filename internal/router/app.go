package router

import (
	"dqix/internal/types"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type app struct {
	dataPath string
	data     data
}

type data struct {
	inventoryMap types.InventoryMap
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
		}

		return nil
	})
}
