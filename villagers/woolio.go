package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	woolio string = "Woolio"
)

var (
	// Woolio is an Animal Crossing villager.
	//
	// Woolio is a Sheep.
	Woolio Villager = villager{
		animal: animals.Sheep.Name(),
		name:   woolio}
)
