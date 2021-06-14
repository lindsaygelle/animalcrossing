package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gayle string = "Gayle"
)

var (
	// Gayle is an Animal Crossing Villager.
	//
	// Gayle is an Alligator.
	Gayle Villager = villager{
		animal: animals.Alligator.Name(),
		name:   gayle}
)
