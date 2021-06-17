package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	phineas string = "Phineas"
)

var (
	// Phineas is an Animal Crossing villager.
	//
	// Phineas is a Fur Seal.
	Phineas Villager = villager{
		animal: animals.FurSeal.Name(),
		name:   phineas}
)
