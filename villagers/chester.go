package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	chester string = "Chester"
)

var (
	// Chester is an Animal Crossing animal.
	//
	// Chester is a Bear.
	Chester Villager = villager{
		animal: animals.Bear.Name(),
		name:   chester}
)
