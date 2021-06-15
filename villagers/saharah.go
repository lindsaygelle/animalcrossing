package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	saharah string = "Saharah"
)

var (
	// Saharah is an Animal Crossing villager.
	//
	// Saharah is a Camel.
	Saharah Villager = villager{
		animal: animals.Camel.Name(),
		name:   saharah}
)
