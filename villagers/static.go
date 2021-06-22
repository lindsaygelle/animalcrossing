package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	static string = "Static"
)

var (
	// Static is an Animal Crossing villager.
	//
	// Static is a Squirrel.
	Static Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   static}
)
