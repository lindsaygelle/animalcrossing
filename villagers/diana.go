package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	diana string = "Diana"
)

var (
	// Diana is an Animal Crossing villager.
	//
	// Diana is a Deer.
	Diana Villager = villager{
		animal: animals.Deer.Name(),
		name:   diana}
)
