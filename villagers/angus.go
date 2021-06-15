package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	angus string = "Angus"
)

var (
	// Angus is an Animal Crossing villager.
	//
	// Angus is a Bull.
	Angus Villager = villager{
		animal: animals.Bull.Name(),
		name:   angus}
)
