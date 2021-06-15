package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	reese string = "Reese"
)

var (
	// Reese is an Animal Crossing villager.
	//
	// Reese is an Alpaca.
	Reese Villager = villager{
		animal: animals.Alpaca.Name(),
		name:   reese}
)
