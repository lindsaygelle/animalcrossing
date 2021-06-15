package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	weldon string = "Weldon"
)

var (
	// Weldon is an Animal Crossing villager.
	//
	// Weldon is a Bull.
	Weldon Villager = villager{
		animal: animals.Bull.Name(),
		name:   weldon}
)
