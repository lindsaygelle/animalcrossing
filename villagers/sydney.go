package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sydney string = "Sydney"
)

var (
	// Sydney is an Animal Crossing villager.
	//
	// Sydney is a Koala.
	Sydney Villager = villager{
		animal: animals.Koala.Name(),
		name:   sydney}
)
