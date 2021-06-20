package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	carmen string = "Carmen"
)

var (
	// Carmen is an Animal Crossing villager.
	//
	// Carmen is a Mouse.
	Carmen Villager = villager{
		animal: animals.Mouse.Name(),
		name:   carmen}
)
