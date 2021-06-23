package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	audie string = "Audie"
)

var (
	// Audie is an Animal Crossing villager.
	//
	// Audie is a Wolf.
	Audie Villager = villager{
		animal: animals.Wolf.Name(),
		name:   audie}
)
