package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tom string = "Tom"
)

var (
	// Tom is an Animal Crossing villager.
	//
	// Tom is a Cat.
	Tom Villager = villager{
		animal: animals.Cat.Name(),
		name:   tom}
)
