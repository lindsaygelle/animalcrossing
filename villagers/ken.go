package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ken string = "Ken"
)

var (
	// Ken is an Animal Crossing villager.
	//
	// Ken is a Chicken.
	Ken Villager = villager{
		animal: animals.Chicken.Name(),
		name:   ken}
)
