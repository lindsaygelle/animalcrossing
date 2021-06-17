package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tucker string = "Tucker"
)

var (
	// Tucker is an Animal Crossing villager.
	//
	// Tucker is an Elephant.
	Tucker Villager = villager{
		animal: animals.Elephant.Name(),
		name:   tucker}
)
