package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pashmina string = "Pashmina"
)

var (
	// Pashmina is an Animal Crossing villager.
	//
	// Pashmina is a Goat.
	Pashmina Villager = villager{
		animal: animals.Goat.Name(),
		name:   pashmina}
)
