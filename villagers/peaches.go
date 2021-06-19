package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	peaches string = "Peaches"
)

var (
	// Peaches is an Animal Crossing villager.
	//
	// Peaches is a Horse.
	Peaches Villager = villager{
		animal: animals.Horse.Name(),
		name:   peaches}
)
