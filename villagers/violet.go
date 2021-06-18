package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	violet string = "Violet"
)

var (
	// Violet is an Animal Crossing villager.
	//
	// Violet is a Gorilla.
	Violet Villager = villager{
		animal: animals.Gorilla.Name(),
		name:   violet}
)
