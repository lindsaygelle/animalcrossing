package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	francine string = "Francine"
)

var (
	// Francine is an Animal Crossing villager.
	//
	// Francine is a Rabbit.
	Francine Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   francine}
)
