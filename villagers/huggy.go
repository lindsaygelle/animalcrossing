package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	huggy string = "Huggy"
)

var (
	// Huggy is an Animal Crossing villager.
	//
	// Huggy is a Koala.
	Huggy Villager = villager{
		animal: animals.Koala.Name(),
		name:   huggy}
)
