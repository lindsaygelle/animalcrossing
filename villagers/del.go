package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	del string = "Del"
)

var (
	// Del is an Animal Crossing Villager.
	//
	// Del is an Alligator.
	Del Villager = villager{
		animal: animals.Alligator.Name(),
		name:   del}
)
