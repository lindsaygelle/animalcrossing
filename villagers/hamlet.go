package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	hamlet string = "Hamlet"
)

var (
	// Hamlet is an Animal Crossing villager.
	//
	// Hamlet is a Hamster.
	Hamlet Villager = villager{
		animal: animals.Hamster.Name(),
		name:   hamlet}
)
