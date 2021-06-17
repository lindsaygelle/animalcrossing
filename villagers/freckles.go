package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	freckles string = "Freckles"
)

var (
	// Freckles is an Animal Crossing villager.
	//
	// Freckles is a Duck.
	Freckles Villager = villager{
		animal: animals.Duck.Name(),
		name:   freckles}
)
