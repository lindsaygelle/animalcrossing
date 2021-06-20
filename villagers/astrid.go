package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	astrid string = "Astrid"
)

var (
	// Astrid is an Animal Crossing villager.
	//
	// Astrid is a Kangaroo.
	Astrid Villager = villager{
		animal: animals.Kangaroo.Name(),
		name:   astrid}
)
