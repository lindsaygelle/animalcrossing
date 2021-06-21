package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	wendy string = "Wendy"
)

var (
	// Wendy is an Animal Crossing villager.
	//
	// Wendy is a Sheep.
	Wendy Villager = villager{
		animal: animals.Sheep.Name(),
		name:   wendy}
)
