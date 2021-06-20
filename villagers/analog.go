package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	analog string = "Analog"
)

var (
	// Analog is an Animal Crossing villager.
	//
	// Analog is a Penguin.
	Analog Villager = villager{
		animal: animals.Penguin.Name(),
		name:   analog}
)
