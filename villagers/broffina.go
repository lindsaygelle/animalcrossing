package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	broffina string = "Broffina"
)

var (
	// Broffina is an Animal Crossing villager.
	//
	// Broffina is a Chicken.
	Broffina Villager = villager{
		animal: animals.Chicken.Name(),
		name:   broffina}
)
