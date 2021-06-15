package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rosie string = "Rosie"
)

var (
	// Rosie is an Animal Crossing villager.
	//
	// Rosie is a Cat.
	Rosie Villager = villager{
		animal: animals.Cat.Name(),
		name:   rosie}
)
