package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	yodel string = "Yodel"
)

var (
	// Yodel is an Animal Crossing villager.
	//
	// Yodel is a Gorilla.
	Yodel Villager = villager{
		animal: animals.Gorilla.Name(),
		name:   yodel}
)
