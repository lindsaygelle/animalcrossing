package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	derwin string = "Derwin"
)

var (
	// Derwin is an Animal Crossing villager.
	//
	// Derwin is a Duck.
	Derwin Villager = villager{
		animal: animals.Duck.Name(),
		name:   derwin}
)
