package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	monique string = "Monique"
)

var (
	// Monique is an Animal Crossing villager.
	//
	// Monique is a Cat.
	Monique Villager = villager{
		animal: animals.Cat.Name(),
		name:   monique}
)
