package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	merengue string = "Merengue"
)

var (
	// Merengue is an Animal Crossing villager.
	//
	// Merengue is a Rhino.
	Merengue Villager = villager{
		animal: animals.Rhinoceros.Name(),
		name:   merengue}
)
