package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ganon string = "Ganon"
)

var (
	// Ganon is an Animal Crossing villager.
	//
	// Ganon is a Pig.
	Ganon Villager = villager{
		animal: animals.Pig.Name(),
		name:   ganon}
)
