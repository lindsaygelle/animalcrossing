package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	june string = "June"
)

var (
	// June is an Animal Crossing animal.
	//
	// June is a Bear.
	June Villager = villager{
		animal: animals.Bear.Name(),
		name:   june}
)
