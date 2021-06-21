package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	hopkins string = "Hopkins"
)

var (
	// Hopkins is an Animal Crossing villager.
	//
	// Hopkins is a Rabbit.
	Hopkins Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   hopkins}
)
