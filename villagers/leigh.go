package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	leigh string = "Leigh"
)

var (
	// Leigh is an Animal Crossing villager.
	//
	// Leigh is a Chicken.
	Leigh Villager = villager{
		animal: animals.Chicken.Name(),
		name:   leigh}
)
