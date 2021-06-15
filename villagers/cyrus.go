package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cyrus string = "Cyrus"
)

var (
	// Cyrus is an Animal Crossing villager.
	//
	// Cyrus is an Alpaca.
	Cyrus Villager = villager{
		animal: animals.Alpaca.Name(),
		name:   cyrus}
)
