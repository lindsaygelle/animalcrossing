package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tabby string = "Tabby"
)

var (
	// Tabby is an Animal Crossing villager.
	//
	// Tabby is a Cat.
	Tabby Villager = villager{
		animal: animals.Cat.Name(),
		name:   tabby}
)
