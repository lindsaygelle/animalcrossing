package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	annalise string = "Annalise"
)

var (
	// Annalise is an Animal Crossing villager.
	//
	// Annalise is a Horse.
	Annalise Villager = villager{
		animal: animals.Horse.Name(),
		name:   annalise}
)
