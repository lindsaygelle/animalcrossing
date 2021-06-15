package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	anchovy string = "Anchovy"
)

var (
	// Anchovy is an Animal Crossing villager.
	//
	// Anchovy is a Bird.
	Anchovy Villager = villager{
		animal: animals.Bird.Name(),
		name:   anchovy}
)
