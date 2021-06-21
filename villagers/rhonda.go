package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rhonda string = "Rhonda"
)

var (
	// Rhonda is an Animal Crossing villager.
	//
	// Rhonda is a Rhino.
	Rhonda Villager = villager{
		animal: animals.Rhinoceros.Name(),
		name:   rhonda}
)
