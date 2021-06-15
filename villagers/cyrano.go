package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cyrano string = "Cyrano"
)

var (
	// Cyrano is an Animal Crossing villager.
	//
	// Cyrano is an Anteater.
	Cyrano Villager = villager{
		animal: animals.Anteater.Name(),
		name:   cyrano}
)
