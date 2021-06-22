package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	franklin string = "Franklin"
)

var (
	// Franklin is an Animal Crossing villager.
	//
	// Franklin is a Turkey.
	Franklin Villager = villager{
		animal: animals.Turkey.Name(),
		name:   franklin}
)
