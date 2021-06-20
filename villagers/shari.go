package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	shari string = "Shari"
)

var (
	// Shari is an Animal Crossing villager.
	//
	// Shari is a Monkey.
	Shari Villager = villager{
		animal: animals.Monkey.Name(),
		name:   shari}
)
