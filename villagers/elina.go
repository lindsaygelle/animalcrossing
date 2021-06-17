package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	elina string = "Elina"
)

var (
	// Elina is an Animal Crossing villager.
	//
	// Elina is an Elephant.
	Elina Villager = villager{
		animal: animals.Elephant.Name(),
		name:   elina}
)
