package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rasher string = "Rasher"
)

var (
	// Rasher is an Animal Crossing villager.
	//
	// Rasher is a Pig.
	Rasher Villager = villager{
		animal: animals.Pig.Name(),
		name:   rasher}
)
