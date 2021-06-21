package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tomNook string = "Tom Nook"
)

var (
	// TomNook is an Animal Crossing villager.
	//
	// TomNook is a Raccoon.
	TomNook Villager = villager{
		animal: animals.Raccoon.Name(),
		name:   tomNook}
)
