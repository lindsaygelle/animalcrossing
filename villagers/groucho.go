package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	groucho string = "Groucho"
)

var (
	// Groucho is an Animal Crossing animal.
	//
	// Groucho is a Bear.
	Groucho Villager = villager{
		animal: animals.Bear.Name(),
		name:   groucho}
)
