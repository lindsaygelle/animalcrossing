package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cranston string = "Cranston"
)

var (
	// Cranston is an Animal Crossing villager.
	//
	// Cranston is an Ostrich.
	Cranston Villager = villager{
		animal: animals.Ostrich.Name(),
		name:   cranston}
)
