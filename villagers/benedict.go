package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	benedict string = "Benedict"
)

var (
	// Benedict is an Animal Crossing villager.
	//
	// Benedict is a Chicken.
	Benedict Villager = villager{
		animal: animals.Chicken.Name(),
		name:   benedict}
)
