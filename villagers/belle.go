package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	belle string = "Belle"
)

var (
	// Belle is an Animal Crossing villager.
	//
	// Belle is a Cow.
	Belle Villager = villager{
		animal: animals.Cow.Name(),
		name:   belle}
)
