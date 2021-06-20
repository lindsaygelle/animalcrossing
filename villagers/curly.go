package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	curly string = "Curly"
)

var (
	// Curly is an Animal Crossing villager.
	//
	// Curly is a Pig.
	Curly Villager = villager{
		animal: animals.Pig.Name(),
		name:   curly}
)
