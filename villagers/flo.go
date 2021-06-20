package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	flo string = "Flo"
)

var (
	// Flo is an Animal Crossing villager.
	//
	// Flo is a Penguin.
	Flo Villager = villager{
		animal: animals.Penguin.Name(),
		name:   flo}
)
