package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	oHare string = "O'Hare"
)

var (
	// OHare is an Animal Crossing villager.
	//
	// OHare is a Rabbit.
	OHare Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   oHare}
)
