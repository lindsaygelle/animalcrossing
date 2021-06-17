package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pompom string = "Pompom"
)

var (
	// Pompom is an Animal Crossing villager.
	//
	// Pompom is a Duck.
	Pompom Villager = villager{
		animal: animals.Duck.Name(),
		name:   pompom}
)
