package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	nana string = "Nana"
)

var (
	// Nana is an Animal Crossing villager.
	//
	// Nana is a Monkey.
	Nana Villager = villager{
		animal: animals.Monkey.Name(),
		name:   nana}
)
