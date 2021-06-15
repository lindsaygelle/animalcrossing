package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cupcake string = "Cupcake"
)

var (
	// Cupcake is an Animal Crossing animal.
	//
	// Cupcake is a Bear.
	Cupcake Villager = villager{
		animal: animals.Bear.Name(),
		name:   cupcake}
)
