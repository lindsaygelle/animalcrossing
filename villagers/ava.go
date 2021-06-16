package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ava string = "Ava"
)

var (
	// Ava is an Animal Crossing villager.
	//
	// Ava is a Chicken.
	Ava Villager = villager{
		animal: animals.Chicken.Name(),
		name:   ava}
)
