package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	lopez string = "Lopez"
)

var (
	// Lopez is an Animal Crossing villager.
	//
	// Lopez is a Deer.
	Lopez Villager = villager{
		animal: animals.Deer.Name(),
		name:   lopez}
)
