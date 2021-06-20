package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	samson string = "Samson"
)

var (
	// Samson is an Animal Crossing villager.
	//
	// Samson is a Mouse.
	Samson Villager = villager{
		animal: animals.Mouse.Name(),
		name:   samson}
)
