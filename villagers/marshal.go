package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	marshal string = "Marshal"
)

var (
	// Marshal is an Animal Crossing villager.
	//
	// Marshal is a Squirrel.
	Marshal Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   marshal}
)
