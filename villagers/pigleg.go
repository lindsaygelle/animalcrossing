package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pigleg string = "Pigleg"
)

var (
	// Pigleg is an Animal Crossing villager.
	//
	// Pigleg is a Pig.
	Pigleg Villager = villager{
		animal: animals.Pig.Name(),
		name:   pigleg}
)
