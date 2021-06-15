package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	jacques string = "Jacques"
)

var (
	// Jacques is an Animal Crossing villager.
	//
	// Jacques is a Bird.
	Jacques Villager = villager{
		animal: animals.Bird.Name(),
		name:   jacques}
)
