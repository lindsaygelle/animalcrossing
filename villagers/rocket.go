package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rocket string = "Rocket"
)

var (
	// Rocket is an Animal Crossing villager.
	//
	// Rocket is a Gorilla.
	Rocket Villager = villager{
		animal: animals.Gorilla.Name(),
		name:   rocket}
)
