package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	grams string = "Grams"
)

var (
	// Grams is an Animal Crossing villager.
	//
	// Grams is a Kappa.
	Grams Villager = villager{
		animal: animals.Kappa.Name(),
		name:   grams}
)
