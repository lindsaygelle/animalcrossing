package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	robin string = "Robin"
)

var (
	// Robin is an Animal Crossing villager.
	//
	// Robin is a Bird.
	Robin Villager = villager{
		animal: animals.Bird.Name(),
		name:   robin}
)
