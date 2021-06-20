package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	lionel string = "Lionel"
)

var (
	// Lionel is an Animal Crossing villager.
	//
	// Lionel is a Lion.
	Lionel Villager = villager{
		animal: animals.Lion.Name(),
		name:   lionel}
)
