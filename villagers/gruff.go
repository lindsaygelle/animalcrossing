package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gruff string = "Gruff"
)

var (
	// Gruff is an Animal Crossing villager.
	//
	// Gruff is a Goat.
	Gruff Villager = villager{
		animal: animals.Goat.Name(),
		name:   gruff}
)
