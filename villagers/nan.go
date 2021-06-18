package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	nan string = "Nan"
)

var (
	// Nan is an Animal Crossing villager.
	//
	// Nan is a Goat.
	Nan Villager = villager{
		animal: animals.Goat.Name(),
		name:   nan}
)
