package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	fruity string = "Fruity"
)

var (
	// Fruity is an Animal Crossing villager.
	//
	// Fruity is a Duck.
	Fruity Villager = villager{
		animal: animals.Duck.Name(),
		name:   fruity}
)
