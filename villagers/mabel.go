package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	mabel string = "Mabel"
)

var (
	// Mabel is an Animal Crossing villager.
	//
	// Mabel is a Hedgehog.
	Mabel Villager = villager{
		animal: animals.Hedgehog.Name(),
		name:   mabel}
)
