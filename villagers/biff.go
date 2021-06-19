package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	biff string = "Biff"
)

var (
	// Biff is an Animal Crossing villager.
	//
	// Biff is a Hippo.
	Biff Villager = villager{
		animal: animals.Hippo.Name(),
		name:   biff}
)
