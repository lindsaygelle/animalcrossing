package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	fauna string = "Fauna"
)

var (
	// Fauna is an Animal Crossing villager.
	//
	// Fauna is a Deer.
	Fauna Villager = villager{
		animal: animals.Deer.Name(),
		name:   fauna}
)
