package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	redd string = "Redd"
)

var (
	// Redd is an Animal Crossing villager.
	//
	// Redd is a Fox.
	Redd Villager = villager{
		animal: animals.Fox.Name(),
		name:   redd}
)
