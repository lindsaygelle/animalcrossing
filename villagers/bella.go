package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bella string = "Bella"
)

var (
	// Bella is an Animal Crossing villager.
	//
	// Bella is a Mouse.
	Bella Villager = villager{
		animal: animals.Mouse.Name(),
		name:   bella}
)
