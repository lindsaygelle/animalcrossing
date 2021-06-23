package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	whitney string = "Whitney"
)

var (
	// Whitney is an Animal Crossing villager.
	//
	// Whitney is a Wolf.
	Whitney Villager = villager{
		animal: animals.Wolf.Name(),
		name:   whitney}
)
