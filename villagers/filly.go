package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	filly string = "Filly"
)

var (
	// Filly is an Animal Crossing villager.
	//
	// Filly is a Horse.
	Filly Villager = villager{
		animal: animals.Horse.Name(),
		name:   filly}
)
