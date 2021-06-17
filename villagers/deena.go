package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	deena string = "Deena"
)

var (
	// Deena is an Animal Crossing villager.
	//
	// Deena is a Duck.
	Deena Villager = villager{
		animal: animals.Duck.Name(),
		name:   deena}
)
