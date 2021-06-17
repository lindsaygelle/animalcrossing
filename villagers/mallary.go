package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	mallary string = "Mallary"
)

var (
	// Mallary is an Animal Crossing villager.
	//
	// Mallary is a Duck.
	Mallary Villager = villager{
		animal: animals.Duck.Name(),
		name:   mallary}
)
