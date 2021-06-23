package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tarou string = "Tarou"
)

var (
	// Tarou is an Animal Crossing villager.
	//
	// Tarou is a Wolf.
	Tarou Villager = villager{
		animal: animals.Wolf.Name(),
		name:   tarou}
)
