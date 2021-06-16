package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rhoda string = "Rhoda"
)

var (
	// Rhoda is an Animal Crossing villager.
	//
	// Rhoda is a Chicken.
	Rhoda Villager = villager{
		animal: animals.Chicken.Name(),
		name:   rhoda}
)
