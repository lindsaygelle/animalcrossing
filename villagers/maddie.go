package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	maddie string = "Maddie"
)

var (
	// Maddie is an Animal Crossing villager.
	//
	// Maddie is a Dog.
	Maddie Villager = villager{
		animal: animals.Dog.Name(),
		name:   maddie}
)
