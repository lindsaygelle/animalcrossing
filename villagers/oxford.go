package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	oxford string = "Oxford"
)

var (
	// Oxford is an Animal Crossing villager.
	//
	// Oxford is a Bull.
	Oxford Villager = villager{
		animal: animals.Bull.Name(),
		name:   oxford}
)
