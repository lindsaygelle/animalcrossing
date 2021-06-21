package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tank string = "Tank"
)

var (
	// Tank is an Animal Crossing villager.
	//
	// Tank is a Rhino.
	Tank Villager = villager{
		animal: animals.Rhinoceros.Name(),
		name:   tank}
)
