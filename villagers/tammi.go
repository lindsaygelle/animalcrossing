package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tammi string = "Tammi"
)

var (
	// Tammi is an Animal Crossing villager.
	//
	// Tammi is a Monkey.
	Tammi Villager = villager{
		animal: animals.Monkey.Name(),
		name:   tammi}
)
