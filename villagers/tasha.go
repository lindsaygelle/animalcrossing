package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tasha string = "Tasha"
)

var (
	// Tasha is an Animal Crossing villager.
	//
	// Tasha is a Squirrel.
	Tasha Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   tasha}
)
