package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gulliver string = "Gulliver"
)

var (
	// Gulliver is an Animal Crossing villager.
	//
	// Gulliver is a Sea Gull.
	Gulliver Villager = villager{
		animal: animals.SeaGull.Name(),
		name:   gulliver}
)
