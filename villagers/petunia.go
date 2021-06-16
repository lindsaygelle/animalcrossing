package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	petunia string = "Petunia"
)

var (
	// Petunia is an Animal Crossing villager.
	//
	// Petunia is a Cow.
	Petunia Villager = villager{
		animal: animals.Cow.Name(),
		name:   petunia}
)
