package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	inkwell string = "Inkwell"
)

var (
	// Inkwell is an Animal Crossing villager.
	//
	// Inkwell is an Octopus.
	Inkwell Villager = villager{
		animal: animals.Octopus.Name(),
		name:   inkwell}
)
