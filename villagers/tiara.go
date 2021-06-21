package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tiara string = "Tiara"
)

var (
	// Tiara is an Animal Crossing villager.
	//
	// Tiara is a Rhino.
	Tiara Villager = villager{
		animal: animals.Rhinoceros.Name(),
		name:   tiara}
)
