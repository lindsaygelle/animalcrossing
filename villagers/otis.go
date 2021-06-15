package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	otis string = "Otis"
)

var (
	// Otis is an Animal Crossing villager.
	//
	// Otis is a Bird.
	Otis Villager = villager{
		animal: animals.Bird.Name(),
		name:   otis}
)
