package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bitty string = "Bitty"
)

var (
	// Bitty is an Animal Crossing villager.
	//
	// Bitty is a Hippo.
	Bitty Villager = villager{
		animal: animals.Hippo.Name(),
		name:   bitty}
)
