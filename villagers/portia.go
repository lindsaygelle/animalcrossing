package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	portia string = "Portia"
)

var (
	// Portia is an Animal Crossing villager.
	//
	// Portia is a Dog.
	Portia Villager = villager{
		animal: animals.Dog.Name(),
		name:   portia}
)
