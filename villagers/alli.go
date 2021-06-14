package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	alli string = "Alli"
)

var (
	// Alli is an Animal Crossing Villager.
	//
	// Alli is an Alligator.
	Alli Villager = villager{
		animal: animals.Alligator.Name(),
		name:   alli}
)
