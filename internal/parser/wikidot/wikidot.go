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
}

func Init(config *parser.Config) parser.Parser {
	return &WikidotParser{
		config:       config,
		inventoryMap: make(types.InventoryMap),
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
		}
	}

	return p.inventoryMap.WriteJSON(filepath.Join(p.config.Path, "inventory"))
}

func (p *WikidotParser) ReadInputFile() (err error) {
	raw, err := os.ReadFile(p.config.Path + p.config.InputFileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, &p.pages)
}
