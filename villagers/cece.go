package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cece string = "Cece"
)

var (
	// Cece is an Animal Crossing villager.
	//
	// Cece is a Squirrel.
	Cece Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   cece}
)
