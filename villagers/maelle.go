package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	maelle string = "Maelle"
)

var (
	// Maelle is an Animal Crossing villager.
	//
	// Maelle is a Duck.
	Maelle Villager = villager{
		animal: animals.Duck.Name(),
		name:   maelle}
)
