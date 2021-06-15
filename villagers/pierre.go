package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pierre string = "Pierre"
)

var (
	// Pierre is an Animal Crossing villager.
	//
	// Pierre is a Cat.
	Pierre Villager = villager{
		animal: animals.Cat.Name(),
		name:   pierre}
)
