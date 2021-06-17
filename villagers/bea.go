package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bea string = "Bea"
)

var (
	// Bea is an Animal Crossing villager.
	//
	// Bea is a Dog.
	Bea Villager = villager{
		animal: animals.Dog.Name(),
		name:   bea}
)
