package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	harriet string = "Harriet"
)

var (
	// Harriet is an Animal Crossing villager.
	//
	// Harriet is a Dog.
	Harriet Villager = villager{
		animal: animals.Dog.Name(),
		name:   harriet}
)
