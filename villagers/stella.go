package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	stella string = "Stella"
)

var (
	// Stella is an Animal Crossing villager.
	//
	// Stella is a Sheep.
	Stella Villager = villager{
		animal: animals.Sheep.Name(),
		name:   stella}
)
