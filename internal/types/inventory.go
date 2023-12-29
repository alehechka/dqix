package types

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type InventoryMap map[string]map[string]map[string]map[string]Inventory

func (i InventoryMap) AddInventory(inventory Inventory) {
	if i == nil {
		return
	}

	if types, ok := i[inventory.Type]; !ok || types == nil {
		i[inventory.Type] = make(map[string]map[string]map[string]Inventory)
	}

	if categories, ok := i[inventory.Type][inventory.Category]; !ok || categories == nil {
		i[inventory.Type][inventory.Category] = make(map[string]map[string]Inventory)
	}

	if classifications, ok := i[inventory.Type][inventory.Category][inventory.Classification]; !ok || classifications == nil {
		i[inventory.Type][inventory.Category][inventory.Classification] = make(map[string]Inventory)
	}

	i[inventory.Type][inventory.Category][inventory.Classification][inventory.ID] = inventory
}

func (i InventoryMap) GetClassification(typeId string, category string, classification string) (classifications map[string]Inventory) {
	types, ok := i[typeId]
	if !ok || types == nil {
		return
	}

	categories, ok := types[category]
	if !ok || categories == nil {
		return
	}

	return categories[classification]
}

func (i InventoryMap) GetClassificationSlice(typeId string, category string, classification string) (classifications []Inventory) {
	classes := i.GetClassification(typeId, category, classification)
	if classes == nil {
		return
	}

	keys := make([]string, 0, len(classes))
	for k := range classes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		classifications = append(classifications, classes[k])
	}
	return
}

func (i InventoryMap) GetInventory(typeId string, category string, classification string, inventoryId string) (inventory Inventory) {
	classifications := i.GetClassification(typeId, category, classification)
	if classifications == nil {
		return
	}

	return classifications[inventoryId]
}

func (i InventoryMap) GetInventoryFromDataKey(d DataKey) Inventory {
	return i.GetInventory(d.Type, d.Category, d.Classification, d.ID)
}

func (i InventoryMap) WriteJSON(basePath string) (err error) {
	for typeId, categories := range i {
		for category, classifications := range categories {
			for classification, inventories := range classifications {
				file, err := json.MarshalIndent(inventories, "", " ")
				if err != nil {
					return err
				}

				path := filepath.Join(basePath, typeId, category)
				if err := os.MkdirAll(path, os.ModePerm); err != nil {
					return err
				}

				filePath := filepath.Join(path, classification+".json")
				if err := os.WriteFile(filePath, file, 0644); err != nil {
					return err
				}
			}
		}
	}
	return
}

type Statistics struct {
	Attack           int     `json:"attack,omitempty"`
	Defense          int     `json:"defense,omitempty"`
	BlockChance      float64 `json:"blockChance,omitempty"`
	Agility          int     `json:"agility,omitempty"`
	EvasionChance    float64 `json:"evasionChance,omitempty"`
	MagicalMight     int     `json:"magicalMight,omitempty"`
	MagicalMending   int     `json:"magicalMending,omitempty"`
	MPAbsorptionRate float64 `json:"mpAbsorptionRate,omitempty"`
	Deftness         int     `json:"deftness,omitempty"`
	Charm            int     `json:"charm,omitempty"`
	Special          struct {
		Usage  string `json:"usage,omitempty"`
		Effect string `json:"effect,omitempty"`
	} `json:"special,omitempty"`
}

type Inventory struct {
	ID             string         `json:"id,omitempty"`
	Title          string         `json:"title,omitempty"`
	Description    string         `json:"description,omitempty"`
	Statistics     Statistics     `json:"statistics,omitempty"`
	Rarity         int            `json:"rarity,omitempty"`
	BuyPrice       int            `json:"buyPrice,omitempty"`
	SellPrice      int            `json:"sellPrice,omitempty"`
	Vocations      []string       `json:"vocations,omitempty"`
	Type           string         `json:"type,omitempty"` // Type can either be `item` or `equipment`
	Category       string         `json:"category,omitempty"`
	Classification string         `json:"classification,omitempty"`
	Recipe         map[string]int `json:"recipe,omitempty"`         // Recipe is a map of ingredients used to alchemize the inventory where the keys are inventory IDs and the values are the number of that inventory needed
	LocationsFound []string       `json:"locationsFound,omitempty"` // LocationsFound represents the locations the Inventory can be found
	DroppedBy      map[string]int `json:"droppedBy,omitempty"`      // DroppedBy is a map of monsters that drop the inventory where the keys are monster IDs and the values are the denominator (x) in the fraction 1/x representing the drop chance
	IngredientFor  []string       `json:"ingredientFor,omitempty"`  // ingredientFor represents the Inventory recipes that this Inventory is part of
	RequiredFor    []string       `json:"requiredFor,omitempty"`
	CanBeUsedFor   []string       `json:"canBeUsedFor,omitempty"`
}

func (i Inventory) GetID() string {
	return i.ID
}

func (i Inventory) GetTitle() string {
	return i.Title
}

func (i Inventory) GetPath() string {
	return "/" + path.Join("inventory", i.Type, i.Category, i.Classification, i.ID)
}

func (i Inventory) ToDataKey() DataKey {
	return DataKey{
		ID:             i.ID,
		Structure:      "inventory",
		Type:           i.Type,
		Category:       i.Category,
		Classification: i.Classification,
		Title:          i.Title,
		Path:           i.GetPath(),
	}
}

func (p PageContent) ParseAsWeapon() (inventory Inventory) {
	inventory.Type = "equipment"
	inventory.Category = "weapon"

	p.parseFromBase(&inventory)
	return
}

func (p PageContent) ParseAsArmor() (inventory Inventory) {
	inventory.Type = "equipment"
	inventory.Category = "armor"

	p.parseFromBase(&inventory)
	return
}

func (p PageContent) ParseAsAccessory() (inventory Inventory) {
	inventory.Type = "equipment"
	inventory.Category = "accessories"

	p.parseFromBase(&inventory)
	return
}

func (p PageContent) ParseAsItem() (inventory Inventory) {
	inventory.Type = "bag"
	inventory.Category = "items"

	p.parseFromBase(&inventory)

	if inventory.Classification == "item" {
		inventory.Classification = "everyday-item"
	}
	return
}

func (p PageContent) parseFromBase(inventory *Inventory) {
	lastIndex := len(p.Text) - 1

	inventory.ID = TitleToID(p.Text[0])
	inventory.Title = p.Text[0]
	inventory.Description = p.Text[1]
	inventory.Classification = p.Text[lastIndex]

	for i := 2; i < lastIndex; i++ {
		switch p.Text[i] {
		// Statistics
		case "Attack:":
			i++
			inventory.Statistics.Agility, _ = strconv.Atoi(p.Text[i])
		case "Defence:":
			i++
			inventory.Statistics.Defense, _ = strconv.Atoi(p.Text[i])
		case "Block Chance:":
			i++
			inventory.Statistics.BlockChance, _ = strconv.ParseFloat(strings.TrimSuffix(p.Text[i], "%"), 64)
		case "Agility:":
			i++
			inventory.Statistics.Agility, _ = strconv.Atoi(p.Text[i])
		case "Evasion Chance:":
			i++
			inventory.Statistics.EvasionChance, _ = strconv.ParseFloat(strings.TrimSuffix(p.Text[i], "%"), 64)
		case "Magical Might:":
			i++
			inventory.Statistics.MagicalMight, _ = strconv.Atoi(p.Text[i])
		case "Magical Mending:":
			i++
			inventory.Statistics.MagicalMending, _ = strconv.Atoi(p.Text[i])
		case "MP Absorption Rate:":
			i++
			inventory.Statistics.MPAbsorptionRate, _ = strconv.ParseFloat(strings.TrimSuffix(p.Text[i], "%"), 64)
		case "Deftness:":
			i++
			inventory.Statistics.Deftness, _ = strconv.Atoi(p.Text[i])
		case "Charm:":
			i++
			inventory.Statistics.Charm, _ = strconv.Atoi(p.Text[i])
		case "Special:":
			stop := i + 2
			for i++; i <= stop; i++ {
				special := p.Text[i]
				if special == "Use:" {
					inventory.Statistics.Special.Usage = p.Text[i+1]
					break
				} else if i < stop {
					inventory.Statistics.Special.Effect = special
					if p.Text[i+1] != "Use:" {
						break
					}
				}
			}
		case "Rarity:":
			i++
			inventory.Rarity, _ = strconv.Atoi(strings.Split(p.Text[i], "/")[0])
		case "Buy price:":
			i++
			inventory.BuyPrice, _ = strconv.Atoi(strings.Split(p.Text[i], " ")[0])
		case "Sell price:":
			i++
			inventory.SellPrice, _ = strconv.Atoi(strings.Split(p.Text[i], " ")[0])
		case "Used by:":
			i++
			rawVocations := strings.Split(p.Text[i], ", ")
			for _, vocation := range rawVocations {
				if strings.HasPrefix(vocation, "All") {
					inventory.Vocations = AllVocations
					break
				}
				inventory.Vocations = append(inventory.Vocations, TitleToID(vocation))
			}
		case "How to make:":
			if inventory.Recipe == nil {
				inventory.Recipe = make(map[string]int)
			}
			for i++; i < lastIndex; i += 2 {
				if !strings.HasPrefix(p.Text[i+1], "x") {
					i--
					break
				}
				id := TitleToID(p.Text[i])
				num, _ := strconv.Atoi(strings.Split(strings.TrimPrefix(p.Text[i+1], "x"), " ")[0])
				inventory.Recipe[id] = num
			}
		case "Where to find:": // TODO
		case "Dropped by:": // TODO
		case "Alchemises:":
			for i++; i < lastIndex; i++ {
				if id := TitleToID(p.Text[i]); id != "" {
					inventory.IngredientFor = append(inventory.IngredientFor, id)
				}
				if i+1 < lastIndex && p.Text[i+1] == "Required for:" {
					break
				}
			}
			sort.Strings(inventory.IngredientFor)
		case "Required for:":
			for i++; i < lastIndex; i++ {
				if id := TitleToID(p.Text[i]); id != "" {
					inventory.RequiredFor = append(inventory.RequiredFor, id)
				}
			}
			sort.Strings(inventory.RequiredFor)
		case "Can be used for:":
			for i++; i < lastIndex; i++ {
				if id := TitleToID(p.Text[i]); id != "" {
					inventory.CanBeUsedFor = append(inventory.CanBeUsedFor, id)
				}
			}
			sort.Strings(inventory.CanBeUsedFor)
		}
	}
}

func TitleToID(title string) string {
	title = strings.TrimSpace(title)
	title = strings.ToLower(title)
	title = strings.TrimPrefix(title, "'")
	title = strings.TrimSuffix(title, "'")
	title = strings.ReplaceAll(title, " ", "-")
	title = strings.ReplaceAll(title, "'", "-")
	title = strings.ReplaceAll(title, ",", "")
	title = strings.ReplaceAll(title, "ä", "ae")

	return title
}
