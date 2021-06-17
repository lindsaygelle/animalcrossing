package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	margie string = "Margie"
)

var (
	// Margie is an Animal Crossing villager.
	//
	// Margie is an Elephant.
	Margie Villager = villager{
		animal: animals.Elephant.Name(),
		name:   margie}
)
