package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	harry string = "Harry"
)

var (
	// Harry is an Animal Crossing villager.
	//
	// Harry is a Hippo.
	Harry Villager = villager{
		animal: animals.Hippo.Name(),
		name:   harry}
)
