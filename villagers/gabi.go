package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gabi string = "Gabi"
)

var (
	// Gabi is an Animal Crossing villager.
	//
	// Gabi is a Rabbit.
	Gabi Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   gabi}
)
