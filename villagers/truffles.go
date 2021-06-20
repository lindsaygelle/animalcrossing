package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	truffles string = "Truffles"
)

var (
	// Truffles is an Animal Crossing villager.
	//
	// Truffles is a Pig.
	Truffles Villager = villager{
		animal: animals.Pig.Name(),
		name:   truffles}
)
