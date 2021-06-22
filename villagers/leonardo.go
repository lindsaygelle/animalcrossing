package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	leonardo string = "Leonardo"
)

var (
	// Leonardo is an Animal Crossing villager.
	//
	// Leonardo is a Tiger.
	Leonardo Villager = villager{
		animal: animals.Tiger.Name(),
		name:   leonardo}
)
