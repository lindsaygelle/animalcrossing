package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	agnes string = "Agnes"
)

var (
	// Agnes is an Animal Crossing villager.
	//
	// Agnes is a Pig.
	Agnes Villager = villager{
		animal: animals.Pig.Name(),
		name:   agnes}
)
