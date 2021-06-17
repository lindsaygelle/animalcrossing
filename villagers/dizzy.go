package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	dizzy string = "Dizzy"
)

var (
	// Dizzy is an Animal Crossing villager.
	//
	// Dizzy is an Elephant.
	Dizzy Villager = villager{
		animal: animals.Elephant.Name(),
		name:   dizzy}
)
