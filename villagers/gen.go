package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gen string = "Gen"
)

var (
	// Gen is an Animal Crossing villager.
	//
	// Gen is a Sheep.
	Gen Villager = villager{
		animal: animals.Sheep.Name(),
		name:   gen}
)
