package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	marcel string = "Marcel"
)

var (
	// Marcel is an Animal Crossing villager.
	//
	// Marcel is a Dog.
	Marcel Villager = villager{
		animal: animals.Dog.Name(),
		name:   marcel}
)
