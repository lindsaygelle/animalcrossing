package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cesar string = "Cesar"
)

var (
	// Cesar is an Animal Crossing villager.
	//
	// Cesar is a Gorilla.
	Cesar Villager = villager{
		animal: animals.Gorilla.Name(),
		name:   cesar}
)
