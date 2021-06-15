package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	joe string = "Joe"
)

var (
	// Joe is an Animal Crossing villager.
	//
	// Joe is a Bird.
	Joe Villager = villager{
		animal: animals.Bird.Name(),
		name:   joe}
)
