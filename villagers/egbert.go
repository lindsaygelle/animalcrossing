package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	egbert string = "Egbert"
)

var (
	// Egbert is an Animal Crossing villager.
	//
	// Egbert is a Chicken.
	Egbert Villager = villager{
		animal: animals.Chicken.Name(),
		name:   egbert}
)
