package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bruce string = "Bruce"
)

var (
	// Bruce is an Animal Crossing villager.
	//
	// Bruce is a Deer.
	Bruce Villager = villager{
		animal: animals.Deer.Name(),
		name:   bruce}
)
