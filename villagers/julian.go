package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	julian string = "Julian"
)

var (
	// Julian is an Animal Crossing villager.
	//
	// Julian is a Horse.
	Julian Villager = villager{
		animal: animals.Horse.Name(),
		name:   julian}
)
