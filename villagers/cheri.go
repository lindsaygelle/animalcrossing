package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cheri string = "Cheri"
)

var (
	// Cheri is an Animal Crossing animal.
	//
	// Cheri is a Bear.
	Cheri Villager = villager{
		animal: animals.Bear.Name(),
		name:   cheri}
)
