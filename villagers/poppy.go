package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	poppy string = "Poppy"
)

var (
	// Poppy is an Animal Crossing villager.
	//
	// Poppy is a Squirrel.
	Poppy Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   poppy}
)
