package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	zell string = "Zell"
)

var (
	// Zell is an Animal Crossing villager.
	//
	// Zell is a Deer.
	Zell Villager = villager{
		animal: animals.Deer.Name(),
		name:   zell}
)
