package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	marina string = "Marina"
)

var (
	// Marina is an Animal Crossing villager.
	//
	// Marina is an Octopus.
	Marina Villager = villager{
		animal: animals.Octopus.Name(),
		name:   marina}
)
