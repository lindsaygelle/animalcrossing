package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	lucky string = "Lucky"
)

var (
	// Lucky is an Animal Crossing villager.
	//
	// Lucky is a Dog.
	Lucky Villager = villager{
		animal: animals.Dog.Name(),
		name:   lucky}
)
