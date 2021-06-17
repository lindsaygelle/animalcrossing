package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	celia string = "Celia"
)

var (
	// Celia is an Animal Crossing villager.
	//
	// Celia is an Eagle.
	Celia Villager = villager{
		animal: animals.Eagle.Name(),
		name:   celia}
)
