package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	goldie string = "Goldie"
)

var (
	// Goldie is an Animal Crossing villager.
	//
	// Goldie is a Dog.
	Goldie Villager = villager{
		animal: animals.Dog.Name(),
		name:   goldie}
)
