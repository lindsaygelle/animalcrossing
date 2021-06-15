package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	jay string = "Jay"
)

var (
	// Jay is an Animal Crossing villager.
	//
	// Jay is a Bird.
	Jay Villager = villager{
		animal: animals.Bird.Name(),
		name:   jay}
)
