package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bunnie string = "Bunnie"
)

var (
	// Bunnie is an Animal Crossing villager.
	//
	// Bunnie is a Rabbit.
	Bunnie Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   bunnie}
)
