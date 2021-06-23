package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	skye string = "Skye"
)

var (
	// Skye is an Animal Crossing villager.
	//
	// Skye is a Wolf.
	Skye Villager = villager{
		animal: animals.Wolf.Name(),
		name:   skye}
)
