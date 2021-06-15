package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rudy string = "Rudy"
)

var (
	// Rudy is an Animal Crossing villager.
	//
	// Rudy is a Cat.
	Rudy Villager = villager{
		animal: animals.Cat.Name(),
		name:   rudy}
)
