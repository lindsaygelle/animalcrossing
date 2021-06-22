package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	luna string = "Luna"
)

var (
	// Luna is an Animal Crossing villager.
	//
	// Luna is a Tapir.
	Luna Villager = villager{
		animal: animals.Tapir.Name(),
		name:   luna}
)
