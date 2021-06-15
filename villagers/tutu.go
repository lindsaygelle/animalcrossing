package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tutu string = "Tutu"
)

var (
	// Tutu is an Animal Crossing animal.
	//
	// Tutu is a Bear.
	Tutu Villager = villager{
		animal: animals.Bear.Name(),
		name:   tutu}
)
