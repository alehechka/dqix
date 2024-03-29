package types

type PageContent struct {
	Path  string
	Text  []string
	Links map[string]string
}

type Pages map[string]PageContent

type DataKey struct {
	Structure      string
	Type           string
	Category       string
	Classification string
	Title          string
	Path           string
}

func (d DataKey) GetID() string {
	return TitleToID(d.Title)
}

func (d DataKey) GetTitle() string {
	return d.Title
}

func (d DataKey) GetPath() string {
	return d.Path
}

const (
	ClassItemEveryday  string = "everyday"
	ClassItemImportant string = "important"

	ClassArmorHead   string = "head"
	ClassArmorTorso  string = "torso"
	ClassArmorShield string = "shield"
	ClassArmorArms   string = "arms"
	ClassArmorLegs   string = "legs"
	ClassArmorFeet   string = "feet"

	ClassWeaponAxe       string = "axes"
	ClassWeaponBoomerang string = "boomerangs"
	ClassWeaponBow       string = "bows"
	ClassWeaponClaw      string = "claws"
	ClassWeaponFan       string = "fans"
	ClassWeaponHammer    string = "hammers"
	ClassWeaponKnife     string = "knives"
	ClassWeaponSpear     string = "spears"
	ClassWeaponStave     string = "staves"
	ClassWeaponSword     string = "swords"
	ClassWeaponWand      string = "wands"
	ClassWeaponWhip      string = "whips"

	ClassAccessories string = "accessories"
)

const (
	FamilyUnknown   = "unknown"
	FamilyAquatic   = "aquatic"
	FamilyBeast     = "beast"
	FamilyBird      = "bird"
	FamilyBug       = "bug"
	FamilyDemon     = "demon"
	FamilyDragon    = "dragon"
	FamilyElemental = "elemental"
	FamilyHumanoid  = "humanoid"
	FamilyMachine   = "machine"
	FamilyMaterial  = "material"
	FamilyPlant     = "plant"
	FamilySlime     = "slime"
	FamilyUndead    = "undead"
	FamilyZombie    = "zombie"
)

func IsWeapon(class string) bool {
	switch class {
	case ClassWeaponAxe,
		ClassWeaponBoomerang,
		ClassWeaponBow,
		ClassWeaponClaw,
		ClassWeaponFan,
		ClassWeaponHammer,
		ClassWeaponKnife,
		ClassWeaponSpear,
		ClassWeaponStave,
		ClassWeaponSword,
		ClassWeaponWand,
		ClassWeaponWhip:
		return true
	default:
		return false
	}
}

func IsArmor(class string) bool {
	switch class {
	case ClassArmorHead,
		ClassArmorTorso,
		ClassArmorShield,
		ClassArmorArms,
		ClassArmorLegs,
		ClassArmorFeet:
		return true
	default:
		return false
	}
}

func IsItem(class string) bool {
	switch class {
	case ClassItemEveryday, ClassItemImportant:
		return true
	default:
		return false
	}
}

func IsMonster(family string) bool {
	switch family {
	case FamilyUnknown,
		FamilyAquatic,
		FamilyBeast,
		FamilyBird,
		FamilyBug,
		FamilyDemon,
		FamilyDragon,
		FamilyElemental,
		FamilyHumanoid,
		FamilyMachine,
		FamilyMaterial,
		FamilyPlant,
		FamilySlime,
		FamilyUndead,
		FamilyZombie:
		return true
	default:
		return false
	}
}
