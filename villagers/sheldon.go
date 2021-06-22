package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sheldon string = "Sheldon"
)

var (
	// Sheldon is an Animal Crossing villager.
	//
	// Sheldon is a Squirrel.
	Sheldon Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   sheldon}
)
