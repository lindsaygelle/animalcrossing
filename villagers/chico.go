package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	chico string = "Chico"
)

var (
	// Chico is an Animal Crossing villager.
	//
	// Chico is a Mouse.
	Chico Villager = villager{
		animal: animals.Mouse.Name(),
		name:   chico}
)
