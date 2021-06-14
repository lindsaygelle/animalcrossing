package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	boots string = "Boots"
)

var (
	// Boots is an Animal Crossing Villager.
	//
	// Boots is an Alligator.
	Boots Villager = villager{
		animal: animals.Alligator.Name(),
		name:   boots}
)
