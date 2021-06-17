package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tia string = "Tia"
)

var (
	// Tia is an Animal Crossing villager.
	//
	// Tia is an Elephant.
	Tia Villager = villager{
		animal: animals.Elephant.Name(),
		name:   tia}
)
