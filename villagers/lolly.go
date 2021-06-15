package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	lolly string = "Lolly"
)

var (
	// Lolly is an Animal Crossing villager.
	//
	// Lolly is a Cat.
	Lolly Villager = villager{
		animal: animals.Cat.Name(),
		name:   lolly}
)
