package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ricky string = "Ricky"
)

var (
	// Ricky is an Animal Crossing villager.
	//
	// Ricky is a Squirrel.
	Ricky Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   ricky}
)
