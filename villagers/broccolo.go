package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	broccolo string = "Broccolo"
)

var (
	// Broccolo is an Animal Crossing villager.
	//
	// Broccolo is a Mouse.
	Broccolo Villager = villager{
		animal: animals.Mouse.Name(),
		name:   broccolo}
)
