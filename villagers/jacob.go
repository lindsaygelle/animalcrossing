package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	jacob string = "Jacob"
)

var (
	// Jacob is an Animal Crossing villager.
	//
	// Jacob is a Bird.
	Jacob Villager = villager{
		animal: animals.Bird.Name(),
		name:   jacob}
)
