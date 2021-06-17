package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	booker string = "Booker"
)

var (
	// Booker is an Animal Crossing villager.
	//
	// Booker is a Dog.
	Booker Villager = villager{
		animal: animals.Dog.Name(),
		name:   booker}
)
