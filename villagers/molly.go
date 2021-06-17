package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	molly string = "Molly"
)

var (
	// Molly is an Animal Crossing villager.
	//
	// Molly is a Duck.
	Molly Villager = villager{
		animal: animals.Duck.Name(),
		name:   molly}
)
