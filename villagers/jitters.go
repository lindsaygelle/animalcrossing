package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	jitters string = "Jitters"
)

var (
	// Jitters is an Animal Crossing villager.
	//
	// Jitters is a Bird.
	Jitters Villager = villager{
		animal: animals.Bird.Name(),
		name:   jitters}
)
