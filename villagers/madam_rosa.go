package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	madamRosa string = "Madam Rosa"
)

var (
	// MadamRosa is an Animal Crossing villager.
	//
	// MadamRosa is a Bird.
	MadamRosa Villager = villager{
		animal: animals.Bird.Name(),
		name:   madamRosa}
)
