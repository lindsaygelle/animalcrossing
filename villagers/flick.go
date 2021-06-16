package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	flick string = "Flick"
)

var (
	// Flick is an Animal Crossing villager.
	//
	// Flick is a Chameleon.
	Flick Villager = villager{
		animal: animals.Chameleon.Name(),
		name:   flick}
)
