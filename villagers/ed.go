package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ed string = "Ed"
)

var (
	// Ed is an Animal Crossing villager.
	//
	// Ed is a Horse.
	Ed Villager = villager{
		animal: animals.Horse.Name(),
		name:   ed}
)
