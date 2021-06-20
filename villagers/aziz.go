package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	aziz string = "Aziz"
)

var (
	// Aziz is an Animal Crossing villager.
	//
	// Aziz is a Lion.
	Aziz Villager = villager{
		animal: animals.Lion.Name(),
		name:   aziz}
)
