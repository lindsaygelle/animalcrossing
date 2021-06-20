package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	valise string = "Valise"
)

var (
	// Valise is an Animal Crossing villager.
	//
	// Valise is a Kangaroo.
	Valise Villager = villager{
		animal: animals.Kangaroo.Name(),
		name:   valise}
)
