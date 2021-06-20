package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rizzo string = "Rizzo"
)

var (
	// Rizzo is an Animal Crossing villager.
	//
	// Rizzo is a Mouse.
	Rizzo Villager = villager{
		animal: animals.Mouse.Name(),
		name:   rizzo}
)
