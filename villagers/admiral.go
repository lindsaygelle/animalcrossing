package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	admiral string = "Admiral"
)

var (
	// Admiral is an Animal Crossing villager.
	//
	// Admiral is a Bird.
	Admiral Villager = villager{
		animal: animals.Bird.Name(),
		name:   admiral}
)
