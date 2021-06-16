package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bessie string = "Bessie"
)

var (
	// Bessie is an Animal Crossing villager.
	//
	// Bessie is a Cow.
	Bessie Villager = villager{
		animal: animals.Cow.Name(),
		name:   bessie}
)
