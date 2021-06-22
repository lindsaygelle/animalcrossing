package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rolf string = "Rolf"
)

var (
	// Rolf is an Animal Crossing villager.
	//
	// Rolf is a Tiger.
	Rolf Villager = villager{
		animal: animals.Tiger.Name(),
		name:   rolf}
)
