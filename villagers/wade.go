package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	wade string = "Wade"
)

var (
	// Wade is an Animal Crossing villager.
	//
	// Wade is a Penguin.
	Wade Villager = villager{
		animal: animals.Penguin.Name(),
		name:   wade}
)
