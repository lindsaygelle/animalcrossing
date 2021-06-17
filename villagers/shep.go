package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	shep string = "Shep"
)

var (
	// Shep is an Animal Crossing villager.
	//
	// Shep is a Dog.
	Shep Villager = villager{
		animal: animals.Dog.Name(),
		name:   shep}
)
