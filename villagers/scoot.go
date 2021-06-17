package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	scoot string = "Scoot"
)

var (
	// Scoot is an Animal Crossing villager.
	//
	// Scoot is a Duck.
	Scoot Villager = villager{
		animal: animals.Duck.Name(),
		name:   scoot}
)
