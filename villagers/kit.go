package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kit string = "Kit"
)

var (
	// Kit is an Animal Crossing villager.
	//
	// Kit is a Squirrel.
	Kit Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   kit}
)
