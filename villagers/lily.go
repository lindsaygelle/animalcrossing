package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	lily string = "Lily"
)

var (
	// Lily is an Animal Crossing villager.
	//
	// Lily is a Frog.
	Lily Villager = villager{
		animal: animals.Frog.Name(),
		name:   lily}
)
