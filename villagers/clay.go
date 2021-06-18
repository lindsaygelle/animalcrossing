package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	clay string = "Clay"
)

var (
	// Clay is an Animal Crossing villager.
	//
	// Clay is a Hamster.
	Clay Villager = villager{
		animal: animals.Hamster.Name(),
		name:   clay}
)
