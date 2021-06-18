package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sherb string = "Sherb"
)

var (
	// Sherb is an Animal Crossing villager.
	//
	// Sherb is a Goat.
	Sherb Villager = villager{
		animal: animals.Goat.Name(),
		name:   sherb}
)
