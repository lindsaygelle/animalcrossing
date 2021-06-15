package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	curt string = "Curt"
)

var (
	// Curt is an Animal Crossing animal.
	//
	// Curt is a Bear.
	Curt Villager = villager{
		animal: animals.Bear.Name(),
		name:   curt}
)
