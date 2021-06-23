package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	dobie string = "Dobie"
)

var (
	// Dobie is an Animal Crossing villager.
	//
	// Dobie is a Wolf.
	Dobie Villager = villager{
		animal: animals.Wolf.Name(),
		name:   dobie}
)
