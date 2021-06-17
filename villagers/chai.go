package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	chai string = "Chai"
)

var (
	// Chai is an Animal Crossing villager.
	//
	// Chai is an Elephant.
	Chai Villager = villager{
		animal: animals.Elephant.Name(),
		name:   chai}
)
