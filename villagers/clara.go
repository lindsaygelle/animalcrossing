package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	clara string = "Clara"
)

var (
	// Clara is an Animal Crossing villager.
	//
	// Clara is a Hippo.
	Clara Villager = villager{
		animal: animals.Hippo.Name(),
		name:   clara}
)
