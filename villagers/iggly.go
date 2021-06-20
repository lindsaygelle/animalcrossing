package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	iggly string = "Iggly"
)

var (
	// Iggly is an Animal Crossing villager.
	//
	// Iggly is a Penguin.
	Iggly Villager = villager{
		animal: animals.Penguin.Name(),
		name:   iggly}
)
