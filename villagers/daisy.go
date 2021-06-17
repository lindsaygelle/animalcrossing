package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	daisy string = "Daisy"
)

var (
	// Daisy is an Animal Crossing villager.
	//
	// Daisy is a Dog.
	Daisy Villager = villager{
		animal: animals.Dog.Name(),
		name:   daisy}
)
