package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sprocket string = "Sprocket"
)

var (
	// Sprocket is an Animal Crossing villager.
	//
	// Sprocket is an Ostrich.
	Sprocket Villager = villager{
		animal: animals.Ostrich.Name(),
		name:   sprocket}
)
