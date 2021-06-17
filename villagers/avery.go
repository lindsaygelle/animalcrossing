package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	avery string = "Avery"
)

var (
	// Avery is an Animal Crossing villager.
	//
	// Avery is an Eagle.
	Avery Villager = villager{
		animal: animals.Eagle.Name(),
		name:   avery}
)
