package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	stitches string = "Stitches"
)

var (
	// Stitches is an Animal Crossing animal.
	//
	// Stitches is a Bear.
	Stitches Villager = villager{
		animal: animals.Bear.Name(),
		name:   stitches}
)
