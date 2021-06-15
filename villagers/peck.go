package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	peck string = "Peck"
)

var (
	// Peck is an Animal Crossing villager.
	//
	// Peck is a Bird.
	Peck Villager = villager{
		animal: animals.Bird.Name(),
		name:   peck}
)
