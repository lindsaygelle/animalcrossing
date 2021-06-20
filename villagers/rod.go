package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rod string = "Rod"
)

var (
	// Rod is an Animal Crossing villager.
	//
	// Rod is a Mouse.
	Rod Villager = villager{
		animal: animals.Mouse.Name(),
		name:   rod}
)
