package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	mac string = "Mac"
)

var (
	// Mac is an Animal Crossing villager.
	//
	// Mac is a Dog.
	Mac Villager = villager{
		animal: animals.Dog.Name(),
		name:   mac}
)
