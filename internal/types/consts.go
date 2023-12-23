package types

const (
	VocationWarrior       string = "warrior"
	VocationMinstrel      string = "minstrel"
	VocationMartialArtist string = "martial-artist"
	VocationGladiator     string = "gladiator"
	VocationPaladin       string = "paladin"
	VocationPriest        string = "priest"
	VocationMage          string = "mage"
	VocationSage          string = "sage"
	VocationLuminary      string = "luminary"
	VocationThief         string = "thief"
	VocationRanger        string = "ranger"
	VocationArmamentalist string = "armamentalist"
)

var AllVocations = []string{
	VocationWarrior, VocationPriest,
	VocationMage, VocationMartialArtist,
	VocationThief, VocationMinstrel,
	VocationGladiator, VocationPaladin,
	VocationArmamentalist, VocationRanger,
	VocationSage, VocationLuminary,
}
