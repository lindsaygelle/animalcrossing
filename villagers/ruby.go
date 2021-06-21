package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ruby string = "Ruby"
)

var (
	// Ruby is an Animal Crossing villager.
	//
	// Ruby is a Rabbit.
	Ruby Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   ruby}
)
