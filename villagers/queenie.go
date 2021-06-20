package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	queenie string = "Queenie"
)

var (
	// Queenie is an Animal Crossing villager.
	//
	// Queenie is an Ostrich.
	Queenie Villager = villager{
		animal: animals.Ostrich.Name(),
		name:   queenie}
)
