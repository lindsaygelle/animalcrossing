package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pecan string = "Pecan"
)

var (
	// Pecan is an Animal Crossing villager.
	//
	// Pecan is a Squirrel.
	Pecan Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   pecan}
)
