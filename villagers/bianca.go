package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bianca string = "Bianca"
)

var (
	// Bianca is an Animal Crossing villager.
	//
	// Bianca is a Tiger.
	Bianca Villager = villager{
		animal: animals.Tiger.Name(),
		name:   bianca}
)
