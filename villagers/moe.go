package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	moe string = "Moe"
)

var (
	// Moe is an Animal Crossing villager.
	//
	// Moe is a Cat.
	Moe Villager = villager{
		animal: animals.Cat.Name(),
		name:   moe}
)
