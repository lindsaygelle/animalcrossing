package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	moose string = "Moose"
)

var (
	// Moose is an Animal Crossing villager.
	//
	// Moose is a Mouse.
	Moose Villager = villager{
		animal: animals.Mouse.Name(),
		name:   moose}
)
