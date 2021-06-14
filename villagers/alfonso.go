package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	alfonso string = "Alfonso"
)

var (
	// Alfonso is an Animal Crossing Villager.
	//
	// Alfonso is an Alligator.
	Alfonso Villager = villager{
		animal: animals.Alligator.Name(),
		name:   alfonso}
)
