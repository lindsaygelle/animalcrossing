package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pierce string = "Pierce"
)

var (
	// Pierce is an Animal Crossing villager.
	//
	// Pierce is an Eagle.
	Pierce Villager = villager{
		animal: animals.Eagle.Name(),
		name:   pierce}
)
