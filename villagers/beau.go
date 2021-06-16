package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	beau string = "Beau"
)

var (
	// Beau is an Animal Crossing villager.
	//
	// Beau is a Deer.
	Beau Villager = villager{
		animal: animals.Deer.Name(),
		name:   beau}
)
