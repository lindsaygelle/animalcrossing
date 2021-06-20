package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	donResetti string = "Don Resetti"
)

var (
	// DonResetti is an Animal Crossing villager.
	//
	// DonResetti is a Mole.
	DonResetti Villager = villager{
		animal: animals.Mole.Name(),
		name:   donResetti}
)
