package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	flossie string = "Flossie"
)

var (
	// Flossie is an Animal Crossing villager.
	//
	// Flossie is a Mouse.
	Flossie Villager = villager{
		animal: animals.Mouse.Name(),
		name:   flossie}
)
