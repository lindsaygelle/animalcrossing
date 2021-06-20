package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	greta string = "Greta"
)

var (
	// Greta is an Animal Crossing villager.
	//
	// Greta is a Mouse.
	Greta Villager = villager{
		animal: animals.Mouse.Name(),
		name:   greta}
)
