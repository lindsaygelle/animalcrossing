package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	simon string = "Simon"
)

var (
	// Simon is an Animal Crossing villager.
	//
	// Simon is a Monkey.
	Simon Villager = villager{
		animal: animals.Monkey.Name(),
		name:   simon}
)
