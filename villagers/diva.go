package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	diva string = "Diva"
)

var (
	// Diva is an Animal Crossing villager.
	//
	// Diva is a Frog.
	Diva Villager = villager{
		animal: animals.Frog.Name(),
		name:   diva}
)
