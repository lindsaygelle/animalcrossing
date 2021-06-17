package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	eloise string = "Eloise"
)

var (
	// Eloise is an Animal Crossing villager.
	//
	// Eloise is an Elephant.
	Eloise Villager = villager{
		animal: animals.Elephant.Name(),
		name:   eloise}
)
