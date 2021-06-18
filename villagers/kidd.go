package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kidd string = "Kidd"
)

var (
	// Kidd is an Animal Crossing villager.
	//
	// Kidd is a Goat.
	Kidd Villager = villager{
		animal: animals.Goat.Name(),
		name:   kidd}
)
