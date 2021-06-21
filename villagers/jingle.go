package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	jingle string = "Jingle"
)

var (
	// Jingle is an Animal Crossing villager.
	//
	// Jingle is a Reindeer.
	Jingle Villager = villager{
		animal: animals.Reindeer.Name(),
		name:   jingle}
)
