package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sally string = "Sally"
)

var (
	// Sally is an Animal Crossing villager.
	//
	// Sally is a Squirrel.
	Sally Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   sally}
)
