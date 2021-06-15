package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	olivia string = "Olivia"
)

var (
	// Olivia is an Animal Crossing villager.
	//
	// Olivia is a Cat.
	Olivia Villager = villager{
		animal: animals.Cat.Name(),
		name:   olivia}
)
