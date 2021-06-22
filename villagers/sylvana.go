package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sylvana string = "Sylvana"
)

var (
	// Sylvana is an Animal Crossing villager.
	//
	// Sylvana is a Squirrel.
	Sylvana Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   sylvana}
)
