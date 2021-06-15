package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bluebear string = "Bluebear"
)

var (
	// Bluebear is an Animal Crossing animal.
	//
	// Bluebear is a Bear.
	Bluebear Villager = villager{
		animal: animals.Bear.Name(),
		name:   bluebear}
)
