package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	boomer string = "Boomer"
)

var (
	// Boomer is an Animal Crossing villager.
	//
	// Boomer is a Penguin.
	Boomer Villager = villager{
		animal: animals.Penguin.Name(),
		name:   boomer}
)
