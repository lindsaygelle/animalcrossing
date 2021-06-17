package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	benjamin string = "Benjamin"
)

var (
	// Benjamin is an Animal Crossing villager.
	//
	// Benjamin is a Dog.
	Benjamin Villager = villager{
		animal: animals.Dog.Name(),
		name:   benjamin}
)
