package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	naomi string = "Naomi"
)

var (
	// Naomi is an Animal Crossing villager.
	//
	// Naomi is a Cow.
	Naomi Villager = villager{
		animal: animals.Cow.Name(),
		name:   naomi}
)
