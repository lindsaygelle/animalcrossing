package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	toby string = "Toby"
)

var (
	// Toby is an Animal Crossing villager.
	//
	// Toby is a Rabbit.
	Toby Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   toby}
)
