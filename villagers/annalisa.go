package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	annalisa string = "Annalisa"
)

var (
	// Annalisa is an Animal Crossing villager.
	//
	// Annalisa is an Anteater.
	Annalisa Villager = villager{
		animal: animals.Anteater.Name(),
		name:   annalisa}
)
