package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	peggy string = "Peggy"
)

var (
	// Peggy is an Animal Crossing villager.
	//
	// Peggy is a Pig.
	Peggy Villager = villager{
		animal: animals.Pig.Name(),
		name:   peggy}
)
