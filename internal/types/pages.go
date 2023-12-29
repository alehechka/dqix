package types

type PageContent struct {
	Path  string
	Text  []string
	Links map[string]string
}

type Pages map[string]PageContent

type DataKey struct {
	ID             string
	Structure      string
	Type           string
	Category       string
	Classification string
}

const (
	ClassItemEveryday  string = "everyday-item"
	ClassItemImportant string = "important-item"

	ClassArmorHead   string = "head"
	ClassArmorTorso  string = "torso"
	ClassArmorShield string = "shield"
	ClassArmorArms   string = "arms"
	ClassArmorLegs   string = "legs"
	ClassArmorFeet   string = "feet"

	ClassWeaponAxe       string = "axe"
	ClassWeaponBoomerang string = "boomerang"
	ClassWeaponBow       string = "bow"
	ClassWeaponClaw      string = "claw"
	ClassWeaponFan       string = "fan"
	ClassWeaponHammer    string = "hammer"
	ClassWeaponKnife     string = "knife"
	ClassWeaponSpear     string = "spear"
	ClassWeaponStave     string = "stave"
	ClassWeaponSword     string = "sword"
	ClassWeaponWand      string = "wand"
	ClassWeaponWhip      string = "whip"
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
