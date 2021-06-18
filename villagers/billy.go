package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	billy string = "Billy"
)

var (
	// Billy is an Animal Crossing villager.
	//
	// Billy is a Goat.
	Billy Villager = villager{
		animal: animals.Goat.Name(),
		name:   billy}
)
