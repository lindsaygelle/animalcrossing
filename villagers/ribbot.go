package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ribbot string = "Ribbot"
)

var (
	// Ribbot is an Animal Crossing villager.
	//
	// Ribbot is a Frog.
	Ribbot Villager = villager{
		animal: animals.Frog.Name(),
		name:   ribbot}
)
