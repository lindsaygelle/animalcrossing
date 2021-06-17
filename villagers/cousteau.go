package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cousteau string = "Cousteau"
)

var (
	// Cousteau is an Animal Crossing villager.
	//
	// Cousteau is a Frog.
	Cousteau Villager = villager{
		animal: animals.Frog.Name(),
		name:   cousteau}
)
