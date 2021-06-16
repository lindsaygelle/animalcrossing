package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	betty string = "Betty"
)

var (
	// Betty is an Animal Crossing villager.
	//
	// Betty is a Chicken.
	Betty Villager = villager{
		animal: animals.Chicken.Name(),
		name:   betty}
)
