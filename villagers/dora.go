package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	dora string = "Dora"
)

var (
	// Dora is an Animal Crossing villager.
	//
	// Dora is a Mouse.
	Dora Villager = villager{
		animal: animals.Mouse.Name(),
		name:   dora}
)
