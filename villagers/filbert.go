package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	filbert string = "Filbert"
)

var (
	// Filbert is an Animal Crossing villager.
	//
	// Filbert is a Squirrel.
	Filbert Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   filbert}
)
