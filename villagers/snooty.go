package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	snooty string = "Snooty"
)

var (
	// Snooty is an Animal Crossing villager.
	//
	// Snooty is an Anteater.
	Snooty Villager = villager{
		animal: animals.Anteater.Name(),
		name:   snooty}
)
