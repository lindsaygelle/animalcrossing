package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	roald string = "Roald"
)

var (
	// Roald is an Animal Crossing villager.
	//
	// Roald is a Penguin.
	Roald Villager = villager{
		animal: animals.Penguin.Name(),
		name:   roald}
)
