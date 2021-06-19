package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	buck string = "Buck"
)

var (
	// Buck is an Animal Crossing villager.
	//
	// Buck is a Horse.
	Buck Villager = villager{
		animal: animals.Horse.Name(),
		name:   buck}
)
