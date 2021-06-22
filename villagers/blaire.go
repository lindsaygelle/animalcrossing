package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	blaire string = "Blaire"
)

var (
	// Blaire is an Animal Crossing villager.
	//
	// Blaire is a Squirrel.
	Blaire Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   blaire}
)
