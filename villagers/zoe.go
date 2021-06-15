package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	zoe string = "Zoe"
)

var (
	// Zoe is an Animal Crossing villager.
	//
	// Zoe is an Anteater.
	Zoe Villager = villager{
		animal: animals.Anteater.Name(),
		name:   zoe}
)
