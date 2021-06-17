package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	prince string = "Prince"
)

var (
	// Prince is an Animal Crossing villager.
	//
	// Prince is a Frog.
	Prince Villager = villager{
		animal: animals.Frog.Name(),
		name:   prince}
)
