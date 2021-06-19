package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	colton string = "Colton"
)

var (
	// Colton is an Animal Crossing villager.
	//
	// Colton is a Horse.
	Colton Villager = villager{
		animal: animals.Horse.Name(),
		name:   colton}
)
