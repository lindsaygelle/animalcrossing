package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	fuchsia string = "Fuchsia"
)

var (
	// Fuchsia is an Animal Crossing villager.
	//
	// Fuchsia is a Deer.
	Fuchsia Villager = villager{
		animal: animals.Deer.Name(),
		name:   fuchsia}
)
