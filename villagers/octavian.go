package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	octavian string = "Octavian"
)

var (
	// Octavian is an Animal Crossing villager.
	//
	// Octavian is an Octopus.
	Octavian Villager = villager{
		animal: animals.Octopus.Name(),
		name:   octavian}
)
