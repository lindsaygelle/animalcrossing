package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	walt string = "Walt"
)

var (
	// Walt is an Animal Crossing villager.
	//
	// Walt is a Kangaroo.
	Walt Villager = villager{
		animal: animals.Kangaroo.Name(),
		name:   walt}
)
