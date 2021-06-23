package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	lobo string = "Lobo"
)

var (
	// Lobo is an Animal Crossing villager.
	//
	// Lobo is a Wolf.
	Lobo Villager = villager{
		animal: animals.Wolf.Name(),
		name:   lobo}
)
