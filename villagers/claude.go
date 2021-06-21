package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	claude string = "Claude"
)

var (
	// Claude is an Animal Crossing villager.
	//
	// Claude is a Rabbit.
	Claude Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   claude}
)
