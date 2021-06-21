package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tommy string = "Tommy"
)

var (
	// Tommy is an Animal Crossing villager.
	//
	// Tommy is a Raccoon.
	Tommy Villager = villager{
		animal: animals.Raccoon.Name(),
		name:   tommy}
)
