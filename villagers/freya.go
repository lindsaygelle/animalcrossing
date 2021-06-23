package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	freya string = "Freya"
)

var (
	// Freya is an Animal Crossing villager.
	//
	// Freya is a Wolf.
	Freya Villager = villager{
		animal: animals.Wolf.Name(),
		name:   freya}
)
