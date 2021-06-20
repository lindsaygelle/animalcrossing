package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	lucy string = "Lucy"
)

var (
	// Lucy is an Animal Crossing villager.
	//
	// Lucy is a Pig.
	Lucy Villager = villager{
		animal: animals.Pig.Name(),
		name:   lucy}
)
