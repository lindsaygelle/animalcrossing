package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	viche string = "Viché"
)

var (
	// Viche is an Animal Crossing villager.
	//
	// Viche is a Squirrel.
	Viche Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   viche}
)
