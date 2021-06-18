package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	jane string = "Jane"
)

var (
	// Jane is an Animal Crossing villager.
	//
	// Jane is a Gorilla.
	Jane Villager = villager{
		animal: animals.Gorilla.Name(),
		name:   jane}
)
