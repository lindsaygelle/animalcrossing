package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	knox string = "Knox"
)

var (
	// Knox is an Animal Crossing villager.
	//
	// Knox is a Chicken.
	Knox Villager = villager{
		animal: animals.Chicken.Name(),
		name:   knox}
)
