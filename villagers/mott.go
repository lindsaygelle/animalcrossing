package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	mott string = "Mott"
)

var (
	// Mott is an Animal Crossing villager.
	//
	// Mott is a Lion.
	Mott Villager = villager{
		animal: animals.Lion.Name(),
		name:   mott}
)
