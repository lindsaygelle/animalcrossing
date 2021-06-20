package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	celeste string = "Celeste"
)

var (
	// Celeste is an Animal Crossing villager.
	//
	// Celeste is an Owl.
	Celeste Villager = villager{
		animal: animals.Owl.Name(),
		name:   celeste}
)
