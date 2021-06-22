package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kicks string = "Kicks"
)

var (
	// Kicks is an Animal Crossing villager.
	//
	// Kicks is a Skunk.
	Kicks Villager = villager{
		animal: animals.Skunk.Name(),
		name:   kicks}
)
