package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	anicotti string = "Anicotti"
)

var (
	// Anicotti is an Animal Crossing villager.
	//
	// Anicotti is a Mouse.
	Anicotti Villager = villager{
		animal: animals.Mouse.Name(),
		name:   anicotti}
)
