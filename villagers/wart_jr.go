package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	wartJr string = "Wart Jr."
)

var (
	// Wart Jr. is an Animal Crossing villager.
	//
	// Wart Jr. is a Frog.
	WartJR Villager = villager{
		animal: animals.Frog.Name(),
		name:   wartJr}
)
