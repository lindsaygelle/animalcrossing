package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bubbles string = "Bubbles"
)

var (
	// Bubbles is an Animal Crossing villager.
	//
	// Bubbles is a Hippo.
	Bubbles Villager = villager{
		animal: animals.Hippo.Name(),
		name:   bubbles}
)
