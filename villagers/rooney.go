package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rooney string = "Rooney"
)

var (
	// Rooney is an Animal Crossing villager.
	//
	// Rooney is a Kangaroo.
	Rooney Villager = villager{
		animal: animals.Kangaroo.Name(),
		name:   rooney}
)
