package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	leopold string = "Leopold"
)

var (
	// Leopold is an Animal Crossing villager.
	//
	// Leopold is a Lion.
	Leopold Villager = villager{
		animal: animals.Lion.Name(),
		name:   leopold}
)
