package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	purrl string = "Purrl"
)

var (
	// Purrl is an Animal Crossing villager.
	//
	// Purrl is a Cat.
	Purrl Villager = villager{
		animal: animals.Cat.Name(),
		name:   purrl}
)
