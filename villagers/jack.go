package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	jack string = "Jack"
)

var (
	// Jack is an Animal Crossing villager.
	//
	// Jack is a Pumpkin.
	Jack Villager = villager{
		animal: animals.Pumpkin.Name(),
		name:   jack}
)
