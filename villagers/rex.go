package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rex string = "Rex"
)

var (
	// Rex is an Animal Crossing villager.
	//
	// Rex is a Lion.
	Rex Villager = villager{
		animal: animals.Lion.Name(),
		name:   rex}
)
