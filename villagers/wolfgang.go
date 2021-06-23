package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	wolfgang string = "Wolfgang"
)

var (
	// Wolfgang is an Animal Crossing villager.
	//
	// Wolfgang is a Wolf.
	Wolfgang Villager = villager{
		animal: animals.Wolf.Name(),
		name:   wolfgang}
)
