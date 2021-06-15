package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rover string = "Rover"
)

var (
	// Rover is an Animal Crossing villager.
	//
	// Rover is a Cat.
	Rover Villager = villager{
		animal: animals.Cat.Name(),
		name:   rover}
)
