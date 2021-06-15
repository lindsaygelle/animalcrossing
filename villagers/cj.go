package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cj string = "C.J."
)

var (
	// CJ is an Animal Crossing villager.
	//
	// CJ is a Beaver.
	CJ Villager = villager{
		animal: animals.Beaver.Name(),
		name:   cj}
)
