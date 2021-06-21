package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tiffany string = "Tiffany"
)

var (
	// Tiffany is an Animal Crossing villager.
	//
	// Tiffany is a Rabbit.
	Tiffany Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   tiffany}
)
