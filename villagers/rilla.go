package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rilla string = "Rilla"
)

var (
	// Rilla is an Animal Crossing villager.
	//
	// Rilla is a Gorilla.
	Rilla Villager = villager{
		animal: animals.Gorilla.Name(),
		name:   rilla}
)
