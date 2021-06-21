package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	patricia string = "Patricia"
)

var (
	// Patricia is an Animal Crossing villager.
	//
	// Patricia is a Rhino.
	Patricia Villager = villager{
		animal: animals.Rhinoceros.Name(),
		name:   patricia}
)
