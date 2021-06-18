package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sven string = "Sven"
)

var (
	// Sven is an Animal Crossing villager.
	//
	// Sven is a Goat.
	Sven Villager = villager{
		animal: animals.Goat.Name(),
		name:   sven}
)
