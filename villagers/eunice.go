package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	eunice string = "Eunice"
)

var (
	// Eunice is an Animal Crossing villager.
	//
	// Eunice is a Sheep.
	Eunice Villager = villager{
		animal: animals.Sheep.Name(),
		name:   eunice}
)
