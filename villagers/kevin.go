package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kevin string = "Kevin"
)

var (
	// Kevin is an Animal Crossing villager.
	//
	// Kevin is a Pig.
	Kevin Villager = villager{
		animal: animals.Pig.Name(),
		name:   kevin}
)
