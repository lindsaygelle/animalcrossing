package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	piper string = "Piper"
)

var (
	// Piper is an Animal Crossing villager.
	//
	// Piper is a Bird.
	Piper Villager = villager{
		animal: animals.Bird.Name(),
		name:   piper}
)
