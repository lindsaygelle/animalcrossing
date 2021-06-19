package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sable string = "Sable"
)

var (
	// Sable is an Animal Crossing villager.
	//
	// Sable is a Hedgehog.
	Sable Villager = villager{
		animal: animals.Hedgehog.Name(),
		name:   sable}
)
