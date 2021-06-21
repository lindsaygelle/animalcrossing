package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	timbra string = "Timbra"
)

var (
	// Timbra is an Animal Crossing villager.
	//
	// Timbra is a Sheep.
	Timbra Villager = villager{
		animal: animals.Sheep.Name(),
		name:   timbra}
)
