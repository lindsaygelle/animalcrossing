package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	chow string = "Chow"
)

var (
	// Chow is an Animal Crossing animal.
	//
	// Chow is a Bear.
	Chow Villager = villager{
		animal: animals.Bear.Name(),
		name:   chow}
)
