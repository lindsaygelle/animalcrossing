package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ketchup string = "Ketchup"
)

var (
	// Ketchup is an Animal Crossing villager.
	//
	// Ketchup is a Duck.
	Ketchup Villager = villager{
		animal: animals.Duck.Name(),
		name:   ketchup}
)
