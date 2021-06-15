package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	marty string = "Marty"
)

var (
	// Marty is an Animal Crossing animal.
	//
	// Marty is a Bear.
	Marty Villager = villager{
		animal: animals.Bear.Name(),
		name:   marty}
)
