package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tortimer string = "Tortimer"
)

var (
	// Tortimer is an Animal Crossing villager.
	//
	// Tortimer is a Tortoise.
	Tortimer Villager = villager{
		animal: animals.Tortoise.Name(),
		name:   tortimer}
)
