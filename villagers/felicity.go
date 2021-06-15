package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	felicity string = "Felicity"
)

var (
	// Felicity is an Animal Crossing villager.
	//
	// Felicity is a Cat.
	Felicity Villager = villager{
		animal: animals.Cat.Name(),
		name:   felicity}
)
