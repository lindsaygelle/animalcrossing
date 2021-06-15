package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	chip string = "Chip"
)

var (
	// Chip is an Animal Crossing villager.
	//
	// Chip is a Beaver.
	Chip Villager = villager{
		animal: animals.Beaver.Name(),
		name:   chip}
)
