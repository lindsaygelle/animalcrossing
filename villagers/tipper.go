package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tipper string = "Tipper"
)

var (
	// Tipper is an Animal Crossing villager.
	//
	// Tipper is a Cow.
	Tipper Villager = villager{
		animal: animals.Cow.Name(),
		name:   tipper}
)
