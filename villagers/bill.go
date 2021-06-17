package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bill string = "Bill"
)

var (
	// Bill is an Animal Crossing villager.
	//
	// Bill is a Duck.
	Bill Villager = villager{
		animal: animals.Duck.Name(),
		name:   bill}
)
