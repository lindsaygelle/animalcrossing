package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	drift string = "Drift"
)

var (
	// Drift is an Animal Crossing villager.
	//
	// Drift is a Frog.
	Drift Villager = villager{
		animal: animals.Frog.Name(),
		name:   drift}
)
