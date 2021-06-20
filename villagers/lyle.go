package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	lyle string = "Lyle"
)

var (
	// Lyle is an Animal Crossing villager.
	//
	// Lyle is an Otter.
	Lyle Villager = villager{
		animal: animals.Otter.Name(),
		name:   lyle}
)
