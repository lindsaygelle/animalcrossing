package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cherry string = "Cherry"
)

var (
	// Cherry is an Animal Crossing villager.
	//
	// Cherry is a Dog.
	Cherry Villager = villager{
		animal: animals.Dog.Name(),
		name:   cherry}
)
