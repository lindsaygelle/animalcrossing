package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pekoe string = "Pekoe"
)

var (
	// Pekoe is an Animal Crossing animal.
	//
	// Pekoe is a Bear.
	Pekoe Villager = villager{
		animal: animals.Bear.Name(),
		name:   pekoe}
)
