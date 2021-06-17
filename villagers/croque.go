package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	croque string = "Croque"
)

var (
	// Croque is an Animal Crossing villager.
	//
	// Croque is a Frog.
	Croque Villager = villager{
		animal: animals.Frog.Name(),
		name:   croque}
)
