package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ellie string = "Ellie"
)

var (
	// Ellie is an Animal Crossing villager.
	//
	// Ellie is an Elephant.
	Ellie Villager = villager{
		animal: animals.Elephant.Name(),
		name:   ellie}
)
