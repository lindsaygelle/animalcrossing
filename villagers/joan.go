package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	joan string = "Joan"
)

var (
	// Joan is an Animal Crossing villager.
	//
	// Joan is a Boar.
	Joan Villager = villager{
		animal: animals.Boar.Name(),
		name:   joan}
)
