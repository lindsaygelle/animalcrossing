package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ursala string = "Ursala"
)

var (
	// Ursala is an Animal Crossing animal.
	//
	// Ursala is a Bear.
	Ursala Villager = villager{
		animal: animals.Bear.Name(),
		name:   ursala}
)
