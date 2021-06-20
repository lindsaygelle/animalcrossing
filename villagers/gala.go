package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gala string = "Gala"
)

var (
	// Gala is an Animal Crossing villager.
	//
	// Gala is a Pig.
	Gala Villager = villager{
		animal: animals.Pig.Name(),
		name:   gala}
)
