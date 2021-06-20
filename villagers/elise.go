package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	elise string = "Elise"
)

var (
	// Elise is an Animal Crossing villager.
	//
	// Elise is a Monkey.
	Elise Villager = villager{
		animal: animals.Monkey.Name(),
		name:   elise}
)
