package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sylvia string = "Sylvia"
)

var (
	// Sylvia is an Animal Crossing villager.
	//
	// Sylvia is a Kangaroo.
	Sylvia Villager = villager{
		animal: animals.Kangaroo.Name(),
		name:   sylvia}
)
