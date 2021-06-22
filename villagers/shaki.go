package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	shaki string = "Shaki"
)

var (
	// Shaki is an Animal Crossing villager.
	//
	// Shaki is a Squirrel.
	Shaki Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   shaki}
)
