package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	hans string = "Hans"
)

var (
	// Hans is an Animal Crossing villager.
	//
	// Hans is a Gorilla.
	Hans Villager = villager{
		animal: animals.Gorilla.Name(),
		name:   hans}
)
