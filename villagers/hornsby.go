package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	hornsby string = "Hornsby"
)

var (
	// Hornsby is an Animal Crossing villager.
	//
	// Hornsby is a Rhino.
	Hornsby Villager = villager{
		animal: animals.Rhinoceros.Name(),
		name:   hornsby}
)
