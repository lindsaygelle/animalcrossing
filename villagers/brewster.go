package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	brewster string = "Brewster"
)

var (
	// Brewster is an Animal Crossing villager.
	//
	// Brewster is a Pigeon.
	Brewster Villager = villager{
		animal: animals.Pigeon.Name(),
		name:   brewster}
)
