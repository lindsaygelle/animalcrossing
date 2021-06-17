package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	raddle string = "Raddle"
)

var (
	// Raddle is an Animal Crossing villager.
	//
	// Raddle is a Frog.
	Raddle Villager = villager{
		animal: animals.Frog.Name(),
		name:   raddle}
)
