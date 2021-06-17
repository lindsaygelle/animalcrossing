package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	frobert string = "Frobert"
)

var (
	// Frobert is an Animal Crossing villager.
	//
	// Frobert is a Frog.
	Frobert Villager = villager{
		animal: animals.Frog.Name(),
		name:   frobert}
)
