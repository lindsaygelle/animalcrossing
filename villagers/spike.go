package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	spike string = "Spike"
)

var (
	// Spike is an Animal Crossing villager.
	//
	// Spike is a Rhino.
	Spike Villager = villager{
		animal: animals.Rhinoceros.Name(),
		name:   spike}
)
