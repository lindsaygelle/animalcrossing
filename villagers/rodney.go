package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rodney string = "Rodney"
)

var (
	// Rodney is an Animal Crossing villager.
	//
	// Rodney is a Hamster.
	Rodney Villager = villager{
		animal: animals.Hamster.Name(),
		name:   rodney}
)
