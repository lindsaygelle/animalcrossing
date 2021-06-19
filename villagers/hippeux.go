package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	hippeux string = "Hippeux"
)

var (
	// Hippeux is an Animal Crossing villager.
	//
	// Hippeux is a Hippo.
	Hippeux Villager = villager{
		animal: animals.Hippo.Name(),
		name:   hippeux}
)
