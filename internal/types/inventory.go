package types

import (
	"cmp"
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type InventoryMap map[string]map[string]map[string]map[string]Inventory

type InventorySlice []Inventory

type HasInventoryStats struct {
	HasAttack           bool
	HasDefense          bool
	HasBlockChance      bool
	HasAgility          bool
	HasEvasionChance    bool
	HasMagicalMight     bool
	HasMagicalMending   bool
	HasMPAbsorptionRate bool
	HasMaxMP            bool
	HasDeftness         bool
	HasCharm            bool
	HasSpecial          bool
}

func (i InventorySlice) GetHasInventoryStats() (stats HasInventoryStats) {
	for _, inventory := range i {
		if inventory.Statistics.Attack > 0 {
			stats.HasAttack = true
		}
		if inventory.Statistics.Defense > 0 {
			stats.HasDefense = true
		}
		if inventory.Statistics.BlockChance > 0 {
			stats.HasBlockChance = true
		}
		if inventory.Statistics.Agility > 0 {
			stats.HasAgility = true
		}
		if inventory.Statistics.EvasionChance > 0 {
			stats.HasEvasionChance = true
		}
		if inventory.Statistics.MagicalMight > 0 {
			stats.HasMagicalMight = true
		}
		if inventory.Statistics.MagicalMending > 0 {
			stats.HasMagicalMending = true
		}
		if inventory.Statistics.MPAbsorptionRate > 0 {
			stats.HasMPAbsorptionRate = true
		}
		if inventory.Statistics.MaxMP > 0 {
			stats.HasMaxMP = true
		}
		if inventory.Statistics.Deftness > 0 {
			stats.HasDeftness = true
		}
		if inventory.Statistics.Charm > 0 {
			stats.HasCharm = true
		}
		if inventory.Statistics.Special.Effect != "" || inventory.Statistics.Special.Usage != "" || inventory.Statistics.Special.Curse != "" {
			stats.HasSpecial = true
		}
	}
	return
}

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

func InventorySortingFunc(sorts Sorts) func(a, b Inventory) int {
	return func(a, b Inventory) int {
		for _, sort := range sorts {
			var comp int
			switch sort.Field {
			case "title", "name":
				comp = cmp.Compare(a.Title, b.Title)
			case "attack":
				comp = cmp.Compare(a.Statistics.Attack, b.Statistics.Attack)
			case "defense":
				comp = cmp.Compare(a.Statistics.Defense, b.Statistics.Defense)
			case "block", "block-chance", "blockChance":
				comp = cmp.Compare(a.Statistics.BlockChance, b.Statistics.BlockChance)
			case "agility":
				comp = cmp.Compare(a.Statistics.Agility, b.Statistics.Agility)
			case "evasion", "evasion-chance", "evasionChance":
				comp = cmp.Compare(a.Statistics.EvasionChance, b.Statistics.EvasionChance)
			case "magical-might", "magicalMight":
				comp = cmp.Compare(a.Statistics.MagicalMight, b.Statistics.MagicalMight)
			case "magical-mending", "magicalMending":
				comp = cmp.Compare(a.Statistics.MagicalMending, b.Statistics.MagicalMending)
			case "mp-absorption-rate", "mp-absorption", "mpAbsorptionRate", "mpAbsorption":
				comp = cmp.Compare(a.Statistics.MPAbsorptionRate, b.Statistics.MPAbsorptionRate)
			case "max-mp", "maxMp", "maxMP":
				comp = cmp.Compare(a.Statistics.MaxMP, b.Statistics.MaxMP)
			case "deftness":
				comp = cmp.Compare(a.Statistics.Deftness, b.Statistics.Deftness)
			case "charm":
				comp = cmp.Compare(a.Statistics.Charm, b.Statistics.Charm)
			}

			if comp == 0 {
				continue
			}
			if sort.Order == SortOrderAsc {
				return comp
			}
			if sort.Order == SortOrderDesc {
				return 0 - comp
			}
		}
		return 0
	}
}

func (i InventoryMap) GetClassificationSlice(typeId string, category string, classification string, sortQuery string) (classifications InventorySlice) {
	classes := i.GetClassification(typeId, category, classification)
	if classes == nil {
		return
	}

	for k := range classes {
		classifications = append(classifications, classes[k])
	}

	sorts := ParseSortingQuery(sortQuery + ",title")
	slices.SortFunc(classifications, InventorySortingFunc(sorts))

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

type Special struct {
	Usage  string `json:"usage,omitempty"`
	Effect string `json:"effect,omitempty"`
	Curse  string `json:"curse,omitempty"`
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
	MaxMP            int     `json:"maxMP,omitempty"`
	Deftness         int     `json:"deftness,omitempty"`
	Charm            int     `json:"charm,omitempty"`
	Special          Special `json:"special,omitempty"`
}

type Inventory struct {
	ID             string            `json:"id,omitempty"`
	Title          string            `json:"title,omitempty"`
	Description    string            `json:"description,omitempty"`
	Statistics     Statistics        `json:"statistics,omitempty"`
	Rarity         int               `json:"rarity,omitempty"`
	BuyPrice       int               `json:"buyPrice,omitempty"`
	SellPrice      int               `json:"sellPrice,omitempty"`
	Vocations      []string          `json:"vocations,omitempty"`
	Type           string            `json:"type,omitempty"` // Type can either be `item` or `equipment`
	Category       string            `json:"category,omitempty"`
	Classification string            `json:"classification,omitempty"`
	Recipe         map[string]int    `json:"recipe,omitempty"`         // Recipe is a map of ingredients used to alchemize the inventory where the keys are inventory IDs and the values are the number of that inventory needed
	LocationsFound []string          `json:"locationsFound,omitempty"` // LocationsFound represents the locations the Inventory can be found
	DroppedBy      map[string]string `json:"droppedBy,omitempty"`      // DroppedBy is a map of monsters that drop the inventory where the keys are monster IDs and the values are the denominator (x) in the fraction 1/x representing the drop chance
	IngredientFor  []string          `json:"ingredientFor,omitempty"`  // ingredientFor represents the Inventory recipes that this Inventory is part of
	RequiredFor    []string          `json:"requiredFor,omitempty"`
	CanBeUsedFor   []string          `json:"canBeUsedFor,omitempty"`
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

func (i Inventory) ImageSrc() string {
	return "/" + path.Join("static", "gallery", i.GetPath()+".png")
}

type Ingredient struct {
	ID       string
	Quantity int
}

func (i Inventory) RecipeSlice() (ingredients []Ingredient) {
	for id, quantity := range i.Recipe {
		ingredients = append(ingredients, Ingredient{ID: id, Quantity: quantity})
	}
	return
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
	inventory.Category = "weapons"

	p.parseFromBase(&inventory)
	return
}

func (p PageContent) ParseAsArmor() (inventory Inventory) {
	inventory.Type = "equipment"
	inventory.Category = "armor"

	p.parseFromBase(&inventory)
	return
}

func (p PageContent) ParseAsItem() (inventory Inventory) {
	inventory.Type = "bag"
	inventory.Category = "items"

	p.parseFromBase(&inventory)

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
			inventory.Statistics.Attack, _ = strconv.Atoi(p.Text[i])
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
		case "Max. MP:":
			i++
			inventory.Statistics.MaxMP, _ = strconv.Atoi(p.Text[i])
		case "Deftness:":
			i++
			inventory.Statistics.Deftness, _ = strconv.Atoi(p.Text[i])
		case "Charm:":
			i++
			inventory.Statistics.Charm, _ = strconv.Atoi(p.Text[i])
		case "Cursed:":
			i++
			inventory.Statistics.Special.Curse = p.Text[i]
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
		case "Where to find:":
			for i++; i < lastIndex; i++ {
				if strings.HasSuffix(p.Text[i], ")") {
					locations := strings.Split(p.Text[i], ", ")
					inventory.LocationsFound = append(inventory.LocationsFound, locations...)
				} else {
					i--
					break
				}
			}
		case "Dropped by:":
			inventory.DroppedBy = make(map[string]string)
			for i++; i < lastIndex; i++ {
				if i+1 < lastIndex && strings.HasPrefix(p.Text[i+1], "(") && (strings.HasSuffix(p.Text[i+1], ")") || strings.HasSuffix(p.Text[i+1], "),")) {
					key := TitleToID(p.Text[i])
					inventory.DroppedBy[key] = strings.TrimSuffix(p.Text[i+1], ",")
				} else {
					i--
					break
				}
			}
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
				if i+1 < lastIndex && p.Text[i+1] == "Can be used for:" {
					break
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
