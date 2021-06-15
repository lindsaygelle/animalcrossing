package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	maple string = "Maple"
)

var (
	// Maple is an Animal Crossing animal.
	//
	// Maple is a Bear.
	Maple Villager = villager{
		animal: animals.Bear.Name(),
		name:   maple}
)
