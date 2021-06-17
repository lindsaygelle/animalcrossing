package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	camofrog string = "Camofrog"
)

var (
	// Camofrog is an Animal Crossing villager.
	//
	// Camofrog is a Frog.
	Camofrog Villager = villager{
		animal: animals.Frog.Name(),
		name:   camofrog}
)
