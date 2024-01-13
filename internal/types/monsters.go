package types

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type MonsterMap map[string]map[string]Monster

func (m MonsterMap) AddMonster(monster Monster) {
	if m == nil {
		return
	}

	if family, ok := m[monster.GetFamilyID()]; !ok || family == nil {
		m[monster.GetFamilyID()] = map[string]Monster{}
	}

	m[monster.GetFamilyID()][monster.GetID()] = monster
}

func (m MonsterMap) WriteJSON(basePath string) (err error) {
	monsterDir := filepath.Join(basePath, "monsters")
	if err := os.MkdirAll(monsterDir, os.ModePerm); err != nil {
		return err
	}

	for familyId, monsters := range m {
		file, err := json.MarshalIndent(monsters, "", " ")
		if err != nil {
			return err
		}

		filePath := filepath.Join(monsterDir, familyId+".json")
		if err := os.WriteFile(filePath, file, 0644); err != nil {
			return err
		}
	}
	return
}

type Monster struct {
	Number               int               `json:"number,omitempty"`
	Title                string            `json:"title,omitempty"`
	Description          string            `json:"description,omitempty"`
	SecondaryDescription string            `json:"secondaryDescription,omitempty"`
	Experience           int               `json:"experience,omitempty"`
	GoldDropped          int               `json:"goldDropped,omitempty"`
	Family               string            `json:"family,omitempty"`
	IsFloating           bool              `json:"isFloating,omitempty"`
	WhereToFind          []string          `json:"whereToFind,omitempty"`
	ItemsDropped         map[string]string `json:"itemsDropped,omitempty"`
	BattleInfo           BattleInfo        `json:"battleInfo,omitempty"`
}

func (m Monster) GetID() string {
	return TitleToID(m.Title)
}

func (m Monster) GetFamilyID() string {
	return TitleToID(m.Family)
}

type BattleInfo struct {
	MaxHP              int            `json:"maxHP,omitempty"`
	MaxMP              int            `json:"maxMP,omitempty"`
	Attack             int            `json:"attack,omitempty"`
	Defense            int            `json:"defense,omitempty"`
	Agility            int            `json:"agility,omitempty"`
	EvasionChance      float64        `json:"evasionChance,omitempty"`
	NumberOfTurns      int            `json:"numberOfTurns,omitempty"`
	ElementalAversions string         `json:"elementalAversions,omitempty"`
	AilmentAffinities  []string       `json:"ailmentAffinities,omitempty"`
	EnragedBy          map[string]int `json:"enragedBy,omitempty"`
}

func (p PageContent) ParseMonster() (monster Monster) {
	lastIndex := len(p.Text) - 1

	rawNumber := p.Text[0][0:3]
	if number, err := strconv.Atoi(rawNumber); err == nil && number > 0 {
		monster.Number = number
		monster.Title = strings.TrimPrefix(p.Text[0], rawNumber+" ")
	} else {
		monster.Title = p.Text[0]
	}

	monster.Description = p.Text[1]
	if p.Text[2] != "Exp.:" {
		monster.SecondaryDescription = p.Text[2]
	}

	for i := 2; i < lastIndex; i++ {
		switch p.Text[i] {
		case "Exp.:":
			i++
			monster.Experience, _ = strconv.Atoi(p.Text[i])
		case "Gold:":
			i++
			monster.GoldDropped, _ = strconv.Atoi(p.Text[i])
		case "Family:":
			i++
			familyParts := strings.Split(p.Text[i], " ")
			monster.Family = familyParts[0]
			if len(familyParts) > 1 && familyParts[1] == "(floating)" {
				monster.IsFloating = true
			}
		case "Where to find:":
			i++
			monster.WhereToFind = strings.Split(p.Text[i], ", ")
		case "Items dropped:":
			monster.ItemsDropped = make(map[string]string)
			for i++; i < lastIndex; i += 2 {
				item := p.Text[i]
				chance := p.Text[i+1]
				monster.ItemsDropped[item] = chance

				if !strings.HasPrefix(p.Text[i+2], "(") {
					break
				}
			}
		case "Max. HP:":
			i++
			monster.BattleInfo.MaxHP, _ = strconv.Atoi(p.Text[i])
		case "Max. MP:":
			i++
			monster.BattleInfo.MaxMP, _ = strconv.Atoi(p.Text[i])
		case "Attack:":
			i++
			monster.BattleInfo.Attack, _ = strconv.Atoi(p.Text[i])
		case "Defence:":
			i++
			monster.BattleInfo.Defense, _ = strconv.Atoi(p.Text[i])
		case "Agility:":
			i++
			monster.GoldDropped, _ = strconv.Atoi(p.Text[i])
		case "Evasion Chance:":
			i++
			monster.BattleInfo.EvasionChance, _ = strconv.ParseFloat(strings.TrimSuffix(p.Text[i], "%"), 64)
		case "# of Turns:":
			i++
			monster.BattleInfo.NumberOfTurns, _ = strconv.Atoi(p.Text[i])
		case "Weak to:", "Elemental Aversions:":
			i++
			monster.BattleInfo.ElementalAversions = p.Text[i]
		case "Ailment Affinities:":
			i++
			monster.BattleInfo.AilmentAffinities = strings.Split(p.Text[i], "; ")
		case "Enraged by:":
			monster.BattleInfo.EnragedBy = make(map[string]int)
			if p.Text[i+1] == "None." {
				continue
			}
			for i++; i < lastIndex; i += 2 {
				ability := p.Text[i]
				chance, _ := strconv.Atoi(strings.TrimPrefix(strings.TrimSuffix(p.Text[i+1], "%)"), "("))
				monster.BattleInfo.EnragedBy[ability] = chance

				if !strings.HasPrefix(p.Text[i+2], "(") {
					break
				}
			}
		}
	}

	return
}
