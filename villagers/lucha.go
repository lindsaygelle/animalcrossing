package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	lucha string = "Lucha"
)

var (
	// Lucha is an Animal Crossing villager.
	//
	// Lucha is a Bird.
	Lucha Villager = villager{
		animal: animals.Bird.Name(),
		name:   lucha}
)
