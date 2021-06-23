package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	chief string = "Chief"
)

var (
	// Chief is an Animal Crossing villager.
	//
	// Chief is a Wolf.
	Chief Villager = villager{
		animal: animals.Wolf.Name(),
		name:   chief}
)
