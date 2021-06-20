package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	leila string = "Leila"
)

var (
	// Leila is an Animal Crossing villager.
	//
	// Leila is a Kappa.
	Leila Villager = villager{
		animal: animals.Kappa.Name(),
		name:   leila}
)
