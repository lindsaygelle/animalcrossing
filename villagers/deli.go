package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	deli string = "Deli"
)

var (
	// Deli is an Animal Crossing villager.
	//
	// Deli is a Monkey.
	Deli Villager = villager{
		animal: animals.Monkey.Name(),
		name:   deli}
)
