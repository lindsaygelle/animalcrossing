package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tybalt string = "Tybalt"
)

var (
	// Tybalt is an Animal Crossing villager.
	//
	// Tybalt is a Tiger.
	Tybalt Villager = villager{
		animal: animals.Tiger.Name(),
		name:   tybalt}
)
