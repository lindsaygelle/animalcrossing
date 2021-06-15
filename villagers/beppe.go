package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	beppe string = "Beppe"
)

var (
	// Beppe is an Animal Crossing villager.
	//
	// Beppe is a Bird.
	Beppe Villager = villager{
		animal: animals.Bird.Name(),
		name:   beppe}
)
