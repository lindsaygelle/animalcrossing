package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	boyd string = "Boyd"
)

var (
	// Boyd is an Animal Crossing villager.
	//
	// Boyd is a Gorilla.
	Boyd Villager = villager{
		animal: animals.Gorilla.Name(),
		name:   boyd}
)
