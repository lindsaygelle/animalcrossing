package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	deirdre string = "Deirdre"
)

var (
	// Deirdre is an Animal Crossing villager.
	//
	// Deirdre is a Deer.
	Deirdre Villager = villager{
		animal: animals.Deer.Name(),
		name:   deirdre}
)
