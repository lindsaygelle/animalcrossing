package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	midge string = "Midge"
)

var (
	// Midge is an Animal Crossing villager.
	//
	// Midge is a Bird.
	Midge Villager = villager{
		animal: animals.Bird.Name(),
		name:   midge}
)
