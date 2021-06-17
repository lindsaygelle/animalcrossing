package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	orville string = "Orville"
)

var (
	// Orville is an Animal Crossing villager.
	//
	// Orville is a Dodo.
	Orville Villager = villager{
		animal: animals.Dodo.Name(),
		name:   orville}
)
