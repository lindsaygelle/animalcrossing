package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	patty string = "Patty"
)

var (
	// Patty is an Animal Crossing villager.
	//
	// Patty is a Cow.
	Patty Villager = villager{
		animal: animals.Cow.Name(),
		name:   patty}
)
