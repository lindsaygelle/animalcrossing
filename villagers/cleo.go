package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cleo string = "Cleo"
)

var (
	// Cleo is an Animal Crossing villager.
	//
	// Cleo is a Horse.
	Cleo Villager = villager{
		animal: animals.Horse.Name(),
		name:   cleo}
)
