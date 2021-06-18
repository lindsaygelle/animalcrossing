package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	chevre string = "Chevre"
)

var (
	// Chevre is an Animal Crossing villager.
	//
	// Chevre is a Goat.
	Chevre Villager = villager{
		animal: animals.Goat.Name(),
		name:   chevre}
)
