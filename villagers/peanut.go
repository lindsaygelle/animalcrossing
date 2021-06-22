package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	peanut string = "Peanut"
)

var (
	// Peanut is an Animal Crossing villager.
	//
	// Peanut is a Squirrel.
	Peanut Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   peanut}
)
