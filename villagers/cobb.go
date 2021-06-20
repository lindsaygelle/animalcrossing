package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cobb string = "Cobb"
)

var (
	// Cobb is an Animal Crossing villager.
	//
	// Cobb is a Pig.
	Cobb Villager = villager{
		animal: animals.Pig.Name(),
		name:   cobb}
)
