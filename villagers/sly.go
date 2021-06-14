package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sly string = "Sly"
)

var (
	// Sly is an Animal Crossing Villager.
	//
	// Sly is an Alligator.
	Sly Villager = villager{
		animal: animals.Alligator.Name(),
		name:   sly}
)
