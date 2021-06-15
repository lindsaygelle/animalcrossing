package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	nate string = "Nate"
)

var (
	// Nate is an Animal Crossing animal.
	//
	// Nate is a Bear.
	Nate Villager = villager{
		animal: animals.Bear.Name(),
		name:   nate}
)
