package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	merry string = "Merry"
)

var (
	// Merry is an Animal Crossing villager.
	//
	// Merry is a Cat.
	Merry Villager = villager{
		animal: animals.Cat.Name(),
		name:   merry}
)
