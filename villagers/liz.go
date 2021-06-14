package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	liz string = "Liz"
)

var (
	// Liz is an Animal Crossing Villager.
	//
	// Liz is an Alligator.
	Liz Villager = villager{
		animal: animals.Alligator.Name(),
		name:   liz}
)
