package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	beardo string = "Beardo"
)

var (
	// Beardo is an Animal Crossing animal.
	//
	// Beardo is a Bear.
	Beardo Villager = villager{
		animal: animals.Bear.Name(),
		name:   beardo}
)
