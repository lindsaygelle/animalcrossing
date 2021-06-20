package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	puck string = "Puck"
)

var (
	// Puck is an Animal Crossing villager.
	//
	// Puck is a Penguin.
	Puck Villager = villager{
		animal: animals.Penguin.Name(),
		name:   puck}
)
