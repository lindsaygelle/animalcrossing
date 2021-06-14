package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	drago string = "Drago"
)

var (
	// Drago is an Animal Crossing Villager.
	//
	// Drago is an Alligator.
	Drago Villager = villager{
		animal: animals.Alligator.Name(),
		name:   drago}
)
