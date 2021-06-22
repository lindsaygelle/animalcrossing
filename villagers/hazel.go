package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	hazel string = "Hazel"
)

var (
	// Hazel is an Animal Crossing villager.
	//
	// Hazel is a Squirrel.
	Hazel Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   hazel}
)
