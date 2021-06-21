package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	timmy string = "Timmy"
)

var (
	// Timmy is an Animal Crossing villager.
	//
	// Timmy is a Raccoon.
	Timmy Villager = villager{
		animal: animals.Raccoon.Name(),
		name:   timmy}
)
