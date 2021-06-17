package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	keaton string = "Keaton"
)

var (
	// Keaton is an Animal Crossing villager.
	//
	// Keaton is an Eagle.
	Keaton Villager = villager{
		animal: animals.Eagle.Name(),
		name:   keaton}
)
