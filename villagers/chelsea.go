package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	chelsea string = "Chelsea"
)

var (
	// Chelsea is an Animal Crossing villager.
	//
	// Chelsea is a Deer.
	Chelsea Villager = villager{
		animal: animals.Deer.Name(),
		name:   chelsea}
)
