package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	genji string = "Genji"
)

var (
	// Genji is an Animal Crossing villager.
	//
	// Genji is a Rabbit.
	Genji Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   genji}
)
