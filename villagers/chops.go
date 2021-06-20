package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	chops string = "Chops"
)

var (
	// Chops is an Animal Crossing villager.
	//
	// Chops is a Pig.
	Chops Villager = villager{
		animal: animals.Pig.Name(),
		name:   chops}
)
