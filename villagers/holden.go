package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	holden string = "Holden"
)

var (
	// Holden is an Animal Crossing villager.
	//
	// Holden is a Hamster.
	Holden Villager = villager{
		animal: animals.Hamster.Name(),
		name:   holden}
)
