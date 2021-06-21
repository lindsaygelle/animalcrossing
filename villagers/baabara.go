package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	baabara string = "Baabara"
)

var (
	// Baabara is an Animal Crossing villager.
	//
	// Baabara is a Sheep.
	Baabara Villager = villager{
		animal: animals.Sheep.Name(),
		name:   baabara}
)
