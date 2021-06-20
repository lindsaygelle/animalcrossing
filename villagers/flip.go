package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	flip string = "Flip"
)

var (
	// Flip is an Animal Crossing villager.
	//
	// Flip is a Monkey.
	Flip Villager = villager{
		animal: animals.Monkey.Name(),
		name:   flip}
)
