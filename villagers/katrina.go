package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	katrina string = "Katrina"
)

var (
	// Katrina is an Animal Crossing villager.
	//
	// Katrina is a Panther.
	Katrina Villager = villager{
		animal: animals.Panther.Name(),
		name:   katrina}
)
