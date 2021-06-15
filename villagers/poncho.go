package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	poncho string = "Poncho"
)

var (
	// Poncho is an Animal Crossing animal.
	//
	// Poncho is a Bear.
	Poncho Villager = villager{
		animal: animals.Bear.Name(),
		name:   poncho}
)
