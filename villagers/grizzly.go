package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	grizzly string = "Grizzly"
)

var (
	// Grizzly is an Animal Crossing animal.
	//
	// Grizzly is a Bear.
	Grizzly Villager = villager{
		animal: animals.Bear.Name(),
		name:   grizzly}
)
