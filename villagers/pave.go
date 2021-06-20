package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pave string = "Pavé"
)

var (
	// Pave is an Animal Crossing villager.
	//
	// Pave is a Peacock.
	Pave Villager = villager{
		animal: animals.Peacock.Name(),
		name:   pave}
)
