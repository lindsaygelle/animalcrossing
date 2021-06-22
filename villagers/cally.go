package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cally string = "Cally"
)

var (
	// Cally is an Animal Crossing villager.
	//
	// Cally is a Squirrel.
	Cally Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   cally}
)
