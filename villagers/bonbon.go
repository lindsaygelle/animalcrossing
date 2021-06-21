package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bonbon string = "Bonbon"
)

var (
	// Bonbon is an Animal Crossing villager.
	//
	// Bonbon is a Rabbit.
	Bonbon Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   bonbon}
)
