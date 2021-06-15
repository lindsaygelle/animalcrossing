package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	antonio string = "Antonio"
)

var (
	// Antonio is an Animal Crossing villager.
	//
	// Antonio is an Anteater.
	Antonio Villager = villager{
		animal: animals.Anteater.Name(),
		name:   antonio}
)
