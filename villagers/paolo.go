package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	paolo string = "Paolo"
)

var (
	// Paolo is an Animal Crossing villager.
	//
	// Paolo is an Elephant.
	Paolo Villager = villager{
		animal: animals.Elephant.Name(),
		name:   paolo}
)
