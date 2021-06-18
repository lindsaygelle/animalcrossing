package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	velma string = "Velma"
)

var (
	// Velma is an Animal Crossing villager.
	//
	// Velma is a Goat.
	Velma Villager = villager{
		animal: animals.Goat.Name(),
		name:   velma}
)
