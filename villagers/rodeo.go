package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rodeo string = "Rodeo"
)

var (
	// Rodeo is an Animal Crossing villager.
	//
	// Rodeo is a Bull.
	Rodeo Villager = villager{
		animal: animals.Bull.Name(),
		name:   rodeo}
)
