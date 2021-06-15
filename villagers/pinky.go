package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pinky string = "Pinky"
)

var (
	// Pinky is an Animal Crossing animal.
	//
	// Pinky is a Bear.
	Pinky Villager = villager{
		animal: animals.Bear.Name(),
		name:   pinky}
)
