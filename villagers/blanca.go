package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	blanca string = "Blanca"
)

var (
	// Blanca is an Animal Crossing villager.
	//
	// Blanca is a Cat.
	Blanca Villager = villager{
		animal: animals.Cat.Name(),
		name:   blanca}
)
