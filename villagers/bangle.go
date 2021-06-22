package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bangle string = "Bangle"
)

var (
	// Bangle is an Animal Crossing villager.
	//
	// Bangle is a Tiger.
	Bangle Villager = villager{
		animal: animals.Tiger.Name(),
		name:   bangle}
)
