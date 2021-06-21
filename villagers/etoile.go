package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	etoile string = "Étoile"
)

var (
	// Etoile is an Animal Crossing villager.
	//
	// Etoile is a Sheep.
	Etoile Villager = villager{
		animal: animals.Sheep.Name(),
		name:   etoile}
)
