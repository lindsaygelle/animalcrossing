package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	iggy string = "Iggy"
)

var (
	// Iggy is an Animal Crossing villager.
	//
	// Iggy is a Goat.
	Iggy Villager = villager{
		animal: animals.Goat.Name(),
		name:   iggy}
)
