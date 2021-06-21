package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	mira string = "Mira"
)

var (
	// Mira is an Animal Crossing villager.
	//
	// Mira is a Rabbit.
	Mira Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   mira}
)
