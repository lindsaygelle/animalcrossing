package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rowan string = "Rowan"
)

var (
	// Rowan is an Animal Crossing villager.
	//
	// Rowan is a Tiger.
	Rowan Villager = villager{
		animal: animals.Tiger.Name(),
		name:   rowan}
)
