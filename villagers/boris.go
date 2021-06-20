package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	boris string = "Boris"
)

var (
	// Boris is an Animal Crossing villager.
	//
	// Boris is a Pig.
	Boris Villager = villager{
		animal: animals.Pig.Name(),
		name:   boris}
)
