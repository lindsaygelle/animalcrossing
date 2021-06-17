package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	drake string = "Drake"
)

var (
	// Drake is an Animal Crossing villager.
	//
	// Drake is a Duck.
	Drake Villager = villager{
		animal: animals.Duck.Name(),
		name:   drake}
)
