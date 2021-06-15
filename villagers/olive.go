package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	olive string = "Olive"
)

var (
	// Olive is an Animal Crossing animal.
	//
	// Olive is a Bear.
	Olive Villager = villager{
		animal: animals.Bear.Name(),
		name:   olive}
)
