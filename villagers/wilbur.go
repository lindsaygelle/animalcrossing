package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	wilbur string = "Wilbur"
)

var (
	// Wilbur is an Animal Crossing villager.
	//
	// Wilbur is a Dodo.
	Wilbur Villager = villager{
		animal: animals.Dodo.Name(),
		name:   wilbur}
)
