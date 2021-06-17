package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	quetzal string = "Quetzal"
)

var (
	// Quetzal is an Animal Crossing villager.
	//
	// Quetzal is an Eagle.
	Quetzal Villager = villager{
		animal: animals.Eagle.Name(),
		name:   quetzal}
)
