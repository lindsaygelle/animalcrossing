package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	harvey string = "Harvey"
)

var (
	// Harvey is an Animal Crossing villager.
	//
	// Harvey is a Dog.
	Harvey Villager = villager{
		animal: animals.Dog.Name(),
		name:   harvey}
)
