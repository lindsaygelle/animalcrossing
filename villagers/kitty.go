package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kitty string = "Kitty"
)

var (
	// Kitty is an Animal Crossing villager.
	//
	// Kitty is a Cat.
	Kitty Villager = villager{
		animal: animals.Cat.Name(),
		name:   kitty}
)
