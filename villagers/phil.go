package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	phil string = "Phil"
)

var (
	// Phil is an Animal Crossing villager.
	//
	// Phil is an Ostrich.
	Phil Villager = villager{
		animal: animals.Ostrich.Name(),
		name:   phil}
)
