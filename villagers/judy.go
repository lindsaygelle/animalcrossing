package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	judy string = "Judy"
)

var (
	// Judy is an Animal Crossing animal.
	//
	// Judy is a Bear.
	Judy Villager = villager{
		animal: animals.Bear.Name(),
		name:   judy}
)
