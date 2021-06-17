package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	huck string = "Huck"
)

var (
	// Huck is an Animal Crossing villager.
	//
	// Huck is a Frog.
	Huck Villager = villager{
		animal: animals.Frog.Name(),
		name:   huck}
)
