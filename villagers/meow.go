package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	meow string = "Meow"
)

var (
	// Meow is an Animal Crossing villager.
	//
	// Meow is a Cat.
	Meow Villager = villager{
		animal: animals.Cat.Name(),
		name:   meow}
)
