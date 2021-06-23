package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	wendell string = "Wendell"
)

var (
	// Wendell is an Animal Crossing villager.
	//
	// Wendell is a Walrus.
	Wendell Villager = villager{
		animal: animals.Walrus.Name(),
		name:   wendell}
)
