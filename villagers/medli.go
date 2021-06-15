package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	medli string = "Medli"
)

var (
	// Medli is an Animal Crossing villager.
	//
	// Medli is a Bird.
	Medli Villager = villager{
		animal: animals.Bird.Name(),
		name:   medli}
)
