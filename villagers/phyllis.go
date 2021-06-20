package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	phyllis string = "Phyllis"
)

var (
	// Phyllis is an Animal Crossing villager.
	//
	// Phyllis is a Pelican.
	Phyllis Villager = villager{
		animal: animals.Pelican.Name(),
		name:   phyllis}
)
