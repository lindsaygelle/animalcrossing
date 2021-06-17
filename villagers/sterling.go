package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sterling string = "Sterling"
)

var (
	// Sterling is an Animal Crossing villager.
	//
	// Sterling is an Eagle.
	Sterling Villager = villager{
		animal: animals.Eagle.Name(),
		name:   sterling}
)
