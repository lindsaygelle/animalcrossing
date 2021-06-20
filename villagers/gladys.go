package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gladys string = "Gladys"
)

var (
	// Gladys is an Animal Crossing villager.
	//
	// Gladys is an Ostrich.
	Gladys Villager = villager{
		animal: animals.Ostrich.Name(),
		name:   gladys}
)
