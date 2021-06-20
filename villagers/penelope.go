package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	penelope string = "Penelope"
)

var (
	// Penelope is an Animal Crossing villager.
	//
	// Penelope is a Mouse.
	Penelope Villager = villager{
		animal: animals.Mouse.Name(),
		name:   penelope}
)
