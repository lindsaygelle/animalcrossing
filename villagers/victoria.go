package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	victoria string = "Victoria"
)

var (
	// Victoria is an Animal Crossing villager.
	//
	// Victoria is a Horse.
	Victoria Villager = villager{
		animal: animals.Horse.Name(),
		name:   victoria}
)
