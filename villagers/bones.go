package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bones string = "Bones"
)

var (
	// Bones is an Animal Crossing villager.
	//
	// Bones is a Dog.
	Bones Villager = villager{
		animal: animals.Dog.Name(),
		name:   bones}
)
