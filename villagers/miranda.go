package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	miranda string = "Miranda"
)

var (
	// Miranda is an Animal Crossing villager.
	//
	// Miranda is a Duck.
	Miranda Villager = villager{
		animal: animals.Duck.Name(),
		name:   miranda}
)
