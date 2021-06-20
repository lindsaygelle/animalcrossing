package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pete string = "Pete"
)

var (
	// Pete is an Animal Crossing villager.
	//
	// Pete is a Pelican.
	Pete Villager = villager{
		animal: animals.Pelican.Name(),
		name:   pete}
)
