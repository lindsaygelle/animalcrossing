package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pascal string = "Pascal"
)

var (
	// Pascal is an Animal Crossing villager.
	//
	// Pascal is an Otter.
	Pascal Villager = villager{
		animal: animals.Otter.Name(),
		name:   pascal}
)
