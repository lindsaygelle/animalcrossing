package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	weber string = "Weber"
)

var (
	// Weber is an Animal Crossing villager.
	//
	// Weber is a Duck.
	Weber Villager = villager{
		animal: animals.Duck.Name(),
		name:   weber}
)
