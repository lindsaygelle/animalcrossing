package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	roscoe string = "Roscoe"
)

var (
	// Roscoe is an Animal Crossing villager.
	//
	// Roscoe is a Horse.
	Roscoe Villager = villager{
		animal: animals.Horse.Name(),
		name:   roscoe}
)
