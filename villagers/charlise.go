package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	charlise string = "Charlise"
)

var (
	// Charlise is an Animal Crossing animal.
	//
	// Charlise is a Bear.
	Charlise Villager = villager{
		animal: animals.Bear.Name(),
		name:   charlise}
)
