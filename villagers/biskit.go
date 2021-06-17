package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	biskit string = "Biskit"
)

var (
	// Biskit is an Animal Crossing villager.
	//
	// Biskit is a Dog.
	Biskit Villager = villager{
		animal: animals.Dog.Name(),
		name:   biskit}
)
