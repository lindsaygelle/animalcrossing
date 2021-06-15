package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	megan string = "Megan"
)

var (
	// Megan is an Animal Crossing animal.
	//
	// Megan is a Bear.
	Megan Villager = villager{
		animal: animals.Bear.Name(),
		name:   megan}
)
