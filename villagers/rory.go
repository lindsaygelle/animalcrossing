package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rory string = "Rory"
)

var (
	// Rory is an Animal Crossing villager.
	//
	// Rory is a Lion.
	Rory Villager = villager{
		animal: animals.Lion.Name(),
		name:   rory}
)
