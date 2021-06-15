package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	punchy string = "Punchy"
)

var (
	// Punchy is an Animal Crossing villager.
	//
	// Punchy is a Cat.
	Punchy Villager = villager{
		animal: animals.Cat.Name(),
		name:   punchy}
)
