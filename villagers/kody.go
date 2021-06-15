package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kody string = "Kody"
)

var (
	// Kody is an Animal Crossing animal.
	//
	// Kody is a Bear.
	Kody Villager = villager{
		animal: animals.Bear.Name(),
		name:   kody}
)
