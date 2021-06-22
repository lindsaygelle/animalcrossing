package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cornimer string = "Cornimer"
)

var (
	// Cornimer is an Animal Crossing villager.
	//
	// Cornimer is a Tortoise.
	Cornimer Villager = villager{
		animal: animals.Tortoise.Name(),
		name:   cornimer}
)
