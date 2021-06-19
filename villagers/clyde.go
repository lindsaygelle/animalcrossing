package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	clyde string = "Clyde"
)

var (
	// Clyde is an Animal Crossing villager.
	//
	// Clyde is a Horse.
	Clyde Villager = villager{
		animal: animals.Horse.Name(),
		name:   clyde}
)
