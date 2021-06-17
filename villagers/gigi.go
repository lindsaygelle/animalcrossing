package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gigi string = "Gigi"
)

var (
	// Gigi is an Animal Crossing villager.
	//
	// Gigi is a Frog.
	Gigi Villager = villager{
		animal: animals.Frog.Name(),
		name:   gigi}
)
