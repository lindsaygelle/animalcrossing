package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sprinkle string = "Sprinkle"
)

var (
	// Sprinkle is an Animal Crossing villager.
	//
	// Sprinkle is a Penguin.
	Sprinkle Villager = villager{
		animal: animals.Penguin.Name(),
		name:   sprinkle}
)
