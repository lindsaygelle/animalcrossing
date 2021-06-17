package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	axel string = "Axel"
)

var (
	// Axel is an Animal Crossing villager.
	//
	// Axel is an Elephant.
	Axel Villager = villager{
		animal: animals.Elephant.Name(),
		name:   axel}
)
