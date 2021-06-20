package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bud string = "Bud"
)

var (
	// Bud is an Animal Crossing villager.
	//
	// Bud is a Lion.
	Bud Villager = villager{
		animal: animals.Lion.Name(),
		name:   bud}
)
