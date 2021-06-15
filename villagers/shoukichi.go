package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	shoukichi string = "Shoukichi"
)

var (
	// Shoukichi is an Animal Crossing villager.
	//
	// Shoukichi is a Bird.
	Shoukichi Villager = villager{
		animal: animals.Bird.Name(),
		name:   shoukichi}
)
