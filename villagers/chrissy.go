package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	chrissy string = "Chrissy"
)

var (
	// Chrissy is an Animal Crossing villager.
	//
	// Chrissy is a Rabbit.
	Chrissy Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   chrissy}
)
