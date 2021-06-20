package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	sueE string = "Sue E"
)

var (
	// SueE is an Animal Crossing villager.
	//
	// SueE is a Pig.
	SueE Villager = villager{
		animal: animals.Pig.Name(),
		name:   sueE}
)
