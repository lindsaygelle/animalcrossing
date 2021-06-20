package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	lyman string = "Lyman"
)

var (
	// Lyman is an Animal Crossing villager.
	//
	// Lyman is a Koala.
	Lyman Villager = villager{
		animal: animals.Koala.Name(),
		name:   lyman}
)
