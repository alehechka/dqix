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
		case "axe", "boomerang", "bow", "claw", "fan", "hammer", "knife", "spear", "stave", "sword", "wand", "whip":
			p.inventoryMap.AddInventory(page.ParseAsWeapon())
		case "important-item", "item":
			p.inventoryMap.AddInventory(page.ParseAsItem())
		case "arms", "head", "feet", "legs", "shield", "torso":
			p.inventoryMap.AddInventory(page.ParseAsArmor())
		case "accessory":
			p.inventoryMap.AddInventory(page.ParseAsAccessory())
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
