package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bigTop string = "Big Top"
)

var (
	// Big Top is an Animal Crossing villager.
	//
	// Big Top is an Elephant.
	BigTop Villager = villager{
		animal: animals.Elephant.Name(),
		name:   bigTop}
)
