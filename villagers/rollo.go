package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rollo string = "Rollo"
)

var (
	// Rollo is an Animal Crossing villager.
	//
	// Rollo is a Hippo.
	Rollo Villager = villager{
		animal: animals.Hippo.Name(),
		name:   rollo}
)
