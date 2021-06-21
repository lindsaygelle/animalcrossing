package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pippy string = "Pippy"
)

var (
	// Pippy is an Animal Crossing villager.
	//
	// Pippy is a Rabbit.
	Pippy Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   pippy}
)
