package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	hopper string = "Hopper"
)

var (
	// Hopper is an Animal Crossing villager.
	//
	// Hopper is a Penguin.
	Hopper Villager = villager{
		animal: animals.Penguin.Name(),
		name:   hopper}
)
