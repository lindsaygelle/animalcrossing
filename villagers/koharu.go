package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	koharu string = "Koharu"
)

var (
	// Koharu is an Animal Crossing villager.
	//
	// Koharu is a Kangaroo.
	Koharu Villager = villager{
		animal: animals.Kangaroo.Name(),
		name:   koharu}
)
