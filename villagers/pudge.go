package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pudge string = "Pudge"
)

var (
	// Pudge is an Animal Crossing animal.
	//
	// Pudge is a Bear.
	Pudge Villager = villager{
		animal: animals.Bear.Name(),
		name:   pudge}
)
