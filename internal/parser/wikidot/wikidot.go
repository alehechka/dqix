package wikidot

import (
	"dqix/internal/parser"
	"dqix/internal/types"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type WikidotParser struct {
	config       *parser.Config
	db           *gorm.DB
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

	if err := p.initDatabase(); err != nil {
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

func (p *WikidotParser) initDatabase() (err error) {
	if _, err := os.Stat(p.config.DatabaseFileName); err == nil {
		if err := os.Remove(p.config.DatabaseFileName); err != nil {
			return err
		}
	}

	if err := os.MkdirAll(strings.TrimSuffix(p.config.DatabaseFileName, filepath.Base(p.config.DatabaseFileName)), os.ModePerm); err != nil {
		return err
	}

	if _, err := os.Create(p.config.DatabaseFileName); err != nil {
		return err
	}

	if p.db, err = gorm.Open(sqlite.Open(p.config.DatabaseFileName), &gorm.Config{}); err != nil {
		return err
	}

	if err := p.db.AutoMigrate(&types.Inventory{}); err != nil {
		return err
	}

	return nil
}
