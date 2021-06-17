package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tad string = "Tad"
)

var (
	// Tad is an Animal Crossing villager.
	//
	// Tad is a Frog.
	Tad Villager = villager{
		animal: animals.Frog.Name(),
		name:   tad}
)
