package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	zucker string = "Zucker"
)

var (
	// Zucker is an Animal Crossing villager.
	//
	// Zucker is an Octopus.
	Zucker Villager = villager{
		animal: animals.Octopus.Name(),
		name:   zucker}
)
