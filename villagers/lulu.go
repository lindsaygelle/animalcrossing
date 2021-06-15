package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	lulu string = "Lulu"
)

var (
	// Lulu is an Animal Crossing villager.
	//
	// Lulu is an Anteater.
	Lulu Villager = villager{
		animal: animals.Anteater.Name(),
		name:   lulu}
)
