package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	marcy string = "Marcy"
)

var (
	// Marcy is an Animal Crossing villager.
	//
	// Marcy is a Kangaroo.
	Marcy Villager = villager{
		animal: animals.Kangaroo.Name(),
		name:   marcy}
)
