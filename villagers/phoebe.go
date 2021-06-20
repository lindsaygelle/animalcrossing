package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	phoebe string = "Phoebe"
)

var (
	// Phoebe is an Animal Crossing villager.
	//
	// Phoebe is an Ostrich.
	Phoebe Villager = villager{
		animal: animals.Ostrich.Name(),
		name:   phoebe}
)
