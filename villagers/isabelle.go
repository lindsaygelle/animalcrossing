package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	isabelle string = "Isabelle"
)

var (
	// Isabelle is an Animal Crossing villager.
	//
	// Isabelle is a Dog.
	Isabelle Villager = villager{
		animal: animals.Dog.Name(),
		name:   isabelle}
)
