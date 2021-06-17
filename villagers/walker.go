package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	walker string = "Walker"
)

var (
	// Walker is an Animal Crossing villager.
	//
	// Walker is a Dog.
	Walker Villager = villager{
		animal: animals.Dog.Name(),
		name:   walker}
)
