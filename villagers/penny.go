package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	penny string = "Penny"
)

var (
	// Penny is an Animal Crossing villager.
	//
	// Penny is a Mouse.
	Penny Villager = villager{
		animal: animals.Mouse.Name(),
		name:   penny}
)
