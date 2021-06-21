package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	dotty string = "Dotty"
)

var (
	// Dotty is an Animal Crossing villager.
	//
	// Dotty is a Rabbit.
	Dotty Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   dotty}
)
