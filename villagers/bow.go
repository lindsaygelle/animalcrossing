package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bow string = "Bow"
)

var (
	// Bow is an Animal Crossing villager.
	//
	// Bow is a Dog.
	Bow Villager = villager{
		animal: animals.Dog.Name(),
		name:   bow}
)
