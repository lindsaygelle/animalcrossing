package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	anabelle string = "Anabelle"
)

var (
	// Anabelle is an Animal Crossing villager.
	//
	// Anabelle is an Anteater.
	Anabelle Villager = villager{
		animal: animals.Anteater.Name(),
		name:   anabelle}
)
