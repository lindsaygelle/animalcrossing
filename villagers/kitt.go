package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kitt string = "Kitt"
)

var (
	// Kitt is an Animal Crossing villager.
	//
	// Kitt is a Kangaroo.
	Kitt Villager = villager{
		animal: animals.Kangaroo.Name(),
		name:   kitt}
)
