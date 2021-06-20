package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	chadder string = "Chadder"
)

var (
	// Chadder is an Animal Crossing villager.
	//
	// Chadder is a Mouse.
	Chadder Villager = villager{
		animal: animals.Mouse.Name(),
		name:   chadder}
)
