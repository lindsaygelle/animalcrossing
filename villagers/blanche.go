package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	blanche string = "Blanche"
)

var (
	// Blanche is an Animal Crossing villager.
	//
	// Blanche is an Ostrich.
	Blanche Villager = villager{
		animal: animals.Ostrich.Name(),
		name:   blanche}
)
