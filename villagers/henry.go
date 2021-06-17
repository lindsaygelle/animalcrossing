package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	henry string = "Henry"
)

var (
	// Henry is an Animal Crossing villager.
	//
	// Henry is a Frog.
	Henry Villager = villager{
		animal: animals.Frog.Name(),
		name:   henry}
)
