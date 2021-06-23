package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kyle string = "Kyle"
)

var (
	// Kyle is an Animal Crossing villager.
	//
	// Kyle is a Wolf.
	Kyle Villager = villager{
		animal: animals.Wolf.Name(),
		name:   kyle}
)
