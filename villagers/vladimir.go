package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	vladimir string = "Vladimir"
)

var (
	// Vladimir is an Animal Crossing animal.
	//
	// Vladimir is a Bear.
	Vladimir Villager = villager{
		animal: animals.Bear.Name(),
		name:   vladimir}
)
