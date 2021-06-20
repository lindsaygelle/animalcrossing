package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	maggie string = "Maggie"
)

var (
	// Maggie is an Animal Crossing villager.
	//
	// Maggie is a Pig.
	Maggie Villager = villager{
		animal: animals.Pig.Name(),
		name:   maggie}
)
