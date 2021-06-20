package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bettina string = "Bettina"
)

var (
	// Bettina is an Animal Crossing villager.
	//
	// Bettina is a Mouse.
	Bettina Villager = villager{
		animal: animals.Mouse.Name(),
		name:   bettina}
)
