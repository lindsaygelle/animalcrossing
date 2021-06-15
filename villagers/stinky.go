package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	stinky string = "Stinky"
)

var (
	// Stinky is an Animal Crossing villager.
	//
	// Stinky is a Cat.
	Stinky Villager = villager{
		animal: animals.Cat.Name(),
		name:   stinky}
)
