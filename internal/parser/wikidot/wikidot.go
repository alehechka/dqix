package wikidot

import (
	"dqix/internal/parser"
	"dqix/internal/types"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type WikidotParser struct {
	config       *parser.Config
	pages        types.Pages
	inventoryMap types.InventoryMap
	monsterMap   types.MonsterMap
}

func Init(config *parser.Config) parser.Parser {
	return &WikidotParser{
		config:       config,
		inventoryMap: make(types.InventoryMap),
		monsterMap:   make(types.MonsterMap),
	}
}

func (p WikidotParser) Parse() (err error) {
	if err := p.ReadInputFile(); err != nil {
		return err
	}

	for path, page := range p.pages {
		if strings.HasPrefix(path, "/system") {
			continue
		}

		if len(page.Text) == 0 {
			continue
		}

		switch page.Text[len(page.Text)-1] {
		case "axes", "boomerangs", "bows", "claws", "fans", "hammers", "knives", "spears", "staves", "swords", "wands", "whips":
			p.inventoryMap.AddInventory(page.ParseAsWeapon())
		case "important", "everyday":
			p.inventoryMap.AddInventory(page.ParseAsItem())
		case "arms", "head", "feet", "legs", "shield", "torso", "accessories":
			p.inventoryMap.AddInventory(page.ParseAsArmor())
		case "monster":
			p.monsterMap.AddMonster(page.ParseMonster())
		}
	}

	if err := os.RemoveAll(p.config.OutputDirectory); err != nil {
		return err
	}

	return p.WriteJSON()
}

func (p *WikidotParser) WriteJSON() (err error) {
	if err := p.inventoryMap.WriteJSON(filepath.Join(p.config.OutputDirectory, "inventory")); err != nil {
		return err
	}

	if err := p.monsterMap.WriteJSON(p.config.OutputDirectory); err != nil {
		return err
	}

	return
}

func (p *WikidotParser) ReadInputFile() (err error) {
	raw, err := os.ReadFile(p.config.Path + p.config.InputFileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, &p.pages)
}
