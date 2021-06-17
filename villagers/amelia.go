package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	amelia string = "Amelia"
)

var (
	// Amelia is an Animal Crossing villager.
	//
	// Amelia is an Eagle.
	Amelia Villager = villager{
		animal: animals.Eagle.Name(),
		name:   amelia}
)
