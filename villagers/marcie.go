package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	marcie string = "Marcie"
)

var (
	// Marcie is an Animal Crossing villager.
	//
	// Marcie is a Kangaroo.
	Marcie Villager = villager{
		animal: animals.Kangaroo.Name(),
		name:   marcie}
)
