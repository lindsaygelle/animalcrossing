package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pironkon string = "Pironkon"
)

var (
	// Pironkon is an Animal Crossing Villager.
	//
	// Pironkon is an Alligator.
	Pironkon Villager = villager{
		animal: animals.Alligator.Name(),
		name:   pironkon}
)
