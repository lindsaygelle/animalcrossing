package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bree string = "Bree"
)

var (
	// Bree is an Animal Crossing villager.
	//
	// Bree is a Mouse.
	Bree Villager = villager{
		animal: animals.Mouse.Name(),
		name:   bree}
)
