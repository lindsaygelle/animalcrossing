package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sunny string = "Sunny"
)

var (
	// Sunny is an Animal Crossing villager.
	//
	// Sunny is a Frog.
	Sunny Villager = villager{
		animal: animals.Frog.Name(),
		name:   sunny}
)
