package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	melba string = "Melba"
)

var (
	// Melba is an Animal Crossing villager.
	//
	// Melba is a Koala.
	Melba Villager = villager{
		animal: animals.Koala.Name(),
		name:   melba}
)
