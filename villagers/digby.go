package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	digby string = "Digby"
)

var (
	// Digby is an Animal Crossing villager.
	//
	// Digby is a Dog.
	Digby Villager = villager{
		animal: animals.Dog.Name(),
		name:   digby}
)
