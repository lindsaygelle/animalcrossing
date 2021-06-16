package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	hank string = "Hank"
)

var (
	// Hank is an Animal Crossing villager.
	//
	// Hank is a Chicken.
	Hank Villager = villager{
		animal: animals.Chicken.Name(),
		name:   hank}
)
