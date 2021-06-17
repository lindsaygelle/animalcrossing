package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	opal string = "Opal"
)

var (
	// Opal is an Animal Crossing villager.
	//
	// Opal is an Elephant.
	Opal Villager = villager{
		animal: animals.Elephant.Name(),
		name:   opal}
)
