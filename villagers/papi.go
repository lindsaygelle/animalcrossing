package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	papi string = "Papi"
)

var (
	// Papi is an Animal Crossing villager.
	//
	// Papi is a Horse.
	Papi Villager = villager{
		animal: animals.Horse.Name(),
		name:   papi}
)
