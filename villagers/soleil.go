package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	soleil string = "Soleil"
)

var (
	// Soleil is an Animal Crossing villager.
	//
	// Soleil is a Hamster.
	Soleil Villager = villager{
		animal: animals.Hamster.Name(),
		name:   soleil}
)
