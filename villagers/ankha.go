package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ankha string = "Ankha"
)

var (
	// Ankha is an Animal Crossing villager.
	//
	// Ankha is a Cat.
	Ankha Villager = villager{
		animal: animals.Cat.Name(),
		name:   ankha}
)
