package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bertha string = "Bertha"
)

var (
	// Bertha is an Animal Crossing villager.
	//
	// Bertha is a Hippo.
	Bertha Villager = villager{
		animal: animals.Hippo.Name(),
		name:   bertha}
)
