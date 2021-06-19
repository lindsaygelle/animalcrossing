package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	elmer string = "Elmer"
)

var (
	// Elmer is an Animal Crossing villager.
	//
	// Elmer is a Horse.
	Elmer Villager = villager{
		animal: animals.Horse.Name(),
		name:   elmer}
)
