package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	carrot string = "Carrot"
)

var (
	// Carrot is an Animal Crossing villager.
	//
	// Carrot is a Cow.
	Carrot Villager = villager{
		animal: animals.Cow.Name(),
		name:   carrot}
)
