package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	al string = "Al"
)

var (
	// Al is an Animal Crossing villager.
	//
	// Al is a Gorilla.
	Al Villager = villager{
		animal: animals.Gorilla.Name(),
		name:   al}
)
