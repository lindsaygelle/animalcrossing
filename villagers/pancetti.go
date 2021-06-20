package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pancetti string = "Pancetti"
)

var (
	// Pancetti is an Animal Crossing villager.
	//
	// Pancetti is a Pig.
	Pancetti Villager = villager{
		animal: animals.Pig.Name(),
		name:   pancetti}
)
