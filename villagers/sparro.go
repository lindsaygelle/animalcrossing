package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sparro string = "Sparro"
)

var (
	// Sparro is an Animal Crossing villager.
	//
	// Sparro is a Bird.
	Sparro Villager = villager{
		animal: animals.Bird.Name(),
		name:   sparro}
)
