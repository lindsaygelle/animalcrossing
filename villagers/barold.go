package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	barold string = "Barold"
)

var (
	// Barold is an Animal Crossing animal.
	//
	// Barold is a Bear.
	Barold Villager = villager{
		animal: animals.Bear.Name(),
		name:   barold}
)
