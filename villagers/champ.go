package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	champ string = "Champ"
)

var (
	// Champ is an Animal Crossing villager.
	//
	// Champ is a Monkey.
	Champ Villager = villager{
		animal: animals.Monkey.Name(),
		name:   champ}
)
