package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	masa string = "Masa"
)

var (
	// Masa is an Animal Crossing villager.
	//
	// Masa is a Dog.
	Masa Villager = villager{
		animal: animals.Dog.Name(),
		name:   masa}
)
