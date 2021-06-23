package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	vivian string = "Vivian"
)

var (
	// Vivian is an Animal Crossing villager.
	//
	// Vivian is a Wolf.
	Vivian Villager = villager{
		animal: animals.Wolf.Name(),
		name:   vivian}
)
