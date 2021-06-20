package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	canberra string = "Canberra"
)

var (
	// Canberra is an Animal Crossing villager.
	//
	// Canberra is a Koala.
	Canberra Villager = villager{
		animal: animals.Koala.Name(),
		name:   canberra}
)
