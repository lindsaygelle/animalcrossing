package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	claudia string = "Claudia"
)

var (
	// Claudia is an Animal Crossing villager.
	//
	// Claudia is a Tiger.
	Claudia Villager = villager{
		animal: animals.Tiger.Name(),
		name:   claudia}
)
