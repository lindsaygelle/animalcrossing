package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	jeremiah string = "Jeremiah"
)

var (
	// Jeremiah is an Animal Crossing villager.
	//
	// Jeremiah is a Frog.
	Jeremiah Villager = villager{
		animal: animals.Frog.Name(),
		name:   jeremiah}
)
