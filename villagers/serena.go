package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	serena string = "Serena"
)

var (
	// Serena is an Animal Crossing villager.
	//
	// Serena is a Dog.
	Serena Villager = villager{
		animal: animals.Dog.Name(),
		name:   serena}
)
