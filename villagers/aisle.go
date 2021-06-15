package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	aisle string = "Aisle"
)

var (
	// Aisle is an Animal Crossing animal.
	//
	// Aisle is a Bear.
	Aisle Villager = villager{
		animal: animals.Bear.Name(),
		name:   aisle}
)
