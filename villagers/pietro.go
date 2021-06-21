package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pietro string = "Pietro"
)

var (
	// Pietro is an Animal Crossing villager.
	//
	// Pietro is a Sheep.
	Pietro Villager = villager{
		animal: animals.Sheep.Name(),
		name:   pietro}
)
