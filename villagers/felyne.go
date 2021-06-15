package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	felyne string = "Felyne"
)

var (
	// Felyne is an Animal Crossing villager.
	//
	// Felyne is a Cat.
	Felyne Villager = villager{
		animal: animals.Cat.Name(),
		name:   felyne}
)
