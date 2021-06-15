package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tangy string = "Tangy"
)

var (
	// Tangy is an Animal Crossing villager.
	//
	// Tangy is a Cat.
	Tangy Villager = villager{
		animal: animals.Cat.Name(),
		name:   tangy}
)
