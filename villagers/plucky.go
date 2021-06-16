package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	plucky string = "Plucky"
)

var (
	// Plucky is an Animal Crossing villager.
	//
	// Plucky is a Chicken.
	Plucky Villager = villager{
		animal: animals.Chicken.Name(),
		name:   plucky}
)
