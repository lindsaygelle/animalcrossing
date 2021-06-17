package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pate string = "Pate"
)

var (
	// Pate is an Animal Crossing villager.
	//
	// Pate is a Duck.
	Pate Villager = villager{
		animal: animals.Duck.Name(),
		name:   pate}
)
