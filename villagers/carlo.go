package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	carlo string = "Carlo"
)

var (
	// Carlo is an Animal Crossing villager.
	//
	// Carlo is a Bird.
	Carlo Villager = villager{
		animal: animals.Bird.Name(),
		name:   carlo}
)
