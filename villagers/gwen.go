package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gwen string = "Gwen"
)

var (
	// Gwen is an Animal Crossing villager.
	//
	// Gwen is a Penguin.
	Gwen Villager = villager{
		animal: animals.Penguin.Name(),
		name:   gwen}
)
