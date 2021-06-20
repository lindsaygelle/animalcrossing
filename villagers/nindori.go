package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	nindori string = "Nindori"
)

var (
	// Nindori is an Animal Crossing villager.
	//
	// Nindori is an Ostrich.
	Nindori Villager = villager{
		animal: animals.Ostrich.Name(),
		name:   nindori}
)
