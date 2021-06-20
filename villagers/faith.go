package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	faith string = "Faith"
)

var (
	// Faith is an Animal Crossing villager.
	//
	// Faith is a Koala.
	Faith Villager = villager{
		animal: animals.Koala.Name(),
		name:   faith}
)
