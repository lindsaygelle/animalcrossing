package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	hector string = "Hector"
)

var (
	// Hector is an Animal Crossing villager.
	//
	// Hector is a Chicken.
	Hector Villager = villager{
		animal: animals.Chicken.Name(),
		name:   hector}
)
