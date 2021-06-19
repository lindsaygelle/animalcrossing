package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	reneigh string = "Reneigh"
)

var (
	// Reneigh is an Animal Crossing villager.
	//
	// Reneigh is a Horse.
	Reneigh Villager = villager{
		animal: animals.Horse.Name(),
		name:   reneigh}
)
