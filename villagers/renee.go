package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	renee string = "Renée"
)

var (
	// Renee is an Animal Crossing villager.
	//
	// Renee is a Rhino.
	Renee Villager = villager{
		animal: animals.Rhinoceros.Name(),
		name:   renee}
)
