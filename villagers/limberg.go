package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	limberg string = "Limberg"
)

var (
	// Limberg is an Animal Crossing villager.
	//
	// Limberg is a Mouse.
	Limberg Villager = villager{
		animal: animals.Mouse.Name(),
		name:   limberg}
)
