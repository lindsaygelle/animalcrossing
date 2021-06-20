package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	julia string = "Julia"
)

var (
	// Julia is an Animal Crossing villager.
	//
	// Julia is an Ostrich.
	Julia Villager = villager{
		animal: animals.Ostrich.Name(),
		name:   julia}
)
