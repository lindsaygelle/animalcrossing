package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	elvis string = "Elvis"
)

var (
	// Elvis is an Animal Crossing villager.
	//
	// Elvis is a Lion.
	Elvis Villager = villager{
		animal: animals.Lion.Name(),
		name:   elvis}
)
