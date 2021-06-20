package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	hugh string = "Hugh"
)

var (
	// Hugh is an Animal Crossing villager.
	//
	// Hugh is a Pig.
	Hugh Villager = villager{
		animal: animals.Pig.Name(),
		name:   hugh}
)
