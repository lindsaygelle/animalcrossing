package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	katie string = "Katie"
)

var (
	// Katie is an Animal Crossing villager.
	//
	// Katie is a Cat.
	Katie Villager = villager{
		animal: animals.Cat.Name(),
		name:   katie}
)
