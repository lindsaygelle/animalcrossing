package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	savannah string = "Savannah"
)

var (
	// Savannah is an Animal Crossing villager.
	//
	// Savannah is a Horse.
	Savannah Villager = villager{
		animal: animals.Horse.Name(),
		name:   savannah}
)
