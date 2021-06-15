package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tammy string = "Tammy"
)

var (
	// Tammy is an Animal Crossing animal.
	//
	// Tammy is a Bear.
	Tammy Villager = villager{
		animal: animals.Bear.Name(),
		name:   tammy}
)
