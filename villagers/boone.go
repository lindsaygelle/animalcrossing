package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	boone string = "Boone"
)

var (
	// Boone is an Animal Crossing villager.
	//
	// Boone is a Gorilla.
	Boone Villager = villager{
		animal: animals.Gorilla.Name(),
		name:   boone}
)
