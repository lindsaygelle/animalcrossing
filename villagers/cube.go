package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cube string = "Cube"
)

var (
	// Cube is an Animal Crossing villager.
	//
	// Cube is a Penguin.
	Cube Villager = villager{
		animal: animals.Penguin.Name(),
		name:   cube}
)
