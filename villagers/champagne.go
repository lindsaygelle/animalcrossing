package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	champagne string = "Champagne"
)

var (
	// Champagne is an Animal Crossing villager.
	//
	// Champagne is a Dog.
	Champagne Villager = villager{
		animal: animals.Dog.Name(),
		name:   champagne}
)
