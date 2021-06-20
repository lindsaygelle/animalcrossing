package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sandy string = "Sandy"
)

var (
	// Sandy is an Animal Crossing villager.
	//
	// Sandy is an Ostrich.
	Sandy Villager = villager{
		animal: animals.Ostrich.Name(),
		name:   sandy}
)
