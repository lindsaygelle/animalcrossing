package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	louie string = "Louie"
)

var (
	// Louie is an Animal Crossing villager.
	//
	// Louie is a Gorilla.
	Louie Villager = villager{
		animal: animals.Gorilla.Name(),
		name:   louie}
)
