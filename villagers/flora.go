package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	flora string = "Flora"
)

var (
	// Flora is an Animal Crossing villager.
	//
	// Flora is an Ostrich.
	Flora Villager = villager{
		animal: animals.Ostrich.Name(),
		name:   flora}
)
