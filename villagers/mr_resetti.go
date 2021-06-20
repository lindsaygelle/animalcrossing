package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	mrResetti string = "Mr. Resetti"
)

var (
	// MRResetti is an Animal Crossing villager.
	//
	// MRResetti is a Mole.
	MRResetti Villager = villager{
		animal: animals.Mole.Name(),
		name:   mrResetti}
)
