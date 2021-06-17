package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	butch string = "Butch"
)

var (
	// Butch is an Animal Crossing villager.
	//
	// Butch is a Dog.
	Butch Villager = villager{
		animal: animals.Dog.Name(),
		name:   butch}
)
