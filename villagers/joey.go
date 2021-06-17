package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	joey string = "Joey"
)

var (
	// Joey is an Animal Crossing villager.
	//
	// Joey is a Duck.
	Joey Villager = villager{
		animal: animals.Duck.Name(),
		name:   joey}
)
