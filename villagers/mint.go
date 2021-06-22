package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	mint string = "Mint"
)

var (
	// Mint is an Animal Crossing villager.
	//
	// Mint is a Squirrel.
	Mint Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   mint}
)
