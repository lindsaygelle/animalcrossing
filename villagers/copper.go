package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	copper string = "Copper"
)

var (
	// Copper is an Animal Crossing villager.
	//
	// Copper is a Dog.
	Copper Villager = villager{
		animal: animals.Dog.Name(),
		name:   copper}
)
