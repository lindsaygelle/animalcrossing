package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pelly string = "Pelly"
)

var (
	// Pelly is an Animal Crossing villager.
	//
	// Pelly is a Pelican.
	Pelly Villager = villager{
		animal: animals.Pelican.Name(),
		name:   pelly}
)
