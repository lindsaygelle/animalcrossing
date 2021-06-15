package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	paula string = "Paula"
)

var (
	// Paula is an Animal Crossing animal.
	//
	// Paula is a Bear.
	Paula Villager = villager{
		animal: animals.Bear.Name(),
		name:   paula}
)
