package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	emerald string = "Emerald"
)

var (
	// Emerald is an Animal Crossing villager.
	//
	// Emerald is a Frog.
	Emerald Villager = villager{
		animal: animals.Frog.Name(),
		name:   emerald}
)
