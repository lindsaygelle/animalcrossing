package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	hambo string = "Hambo"
)

var (
	// Hambo is an Animal Crossing villager.
	//
	// Hambo is a Pig.
	Hambo Villager = villager{
		animal: animals.Pig.Name(),
		name:   hambo}
)
