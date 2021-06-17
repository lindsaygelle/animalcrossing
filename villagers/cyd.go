package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cyd string = "Cyd"
)

var (
	// Cyd is an Animal Crossing villager.
	//
	// Cyd is an Elephant.
	Cyd Villager = villager{
		animal: animals.Elephant.Name(),
		name:   cyd}
)
