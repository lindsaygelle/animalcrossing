package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	fang string = "Fang"
)

var (
	// Fang is an Animal Crossing villager.
	//
	// Fang is a Wolf.
	Fang Villager = villager{
		animal: animals.Wolf.Name(),
		name:   fang}
)
