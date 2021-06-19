package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rocco string = "Rocco"
)

var (
	// Rocco is an Animal Crossing villager.
	//
	// Rocco is a Hippo.
	Rocco Villager = villager{
		animal: animals.Hippo.Name(),
		name:   rocco}
)
