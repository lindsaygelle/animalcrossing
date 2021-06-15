package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	vic string = "Vic"
)

var (
	// Vic is an Animal Crossing villager.
	//
	// Vic is a Bull.
	Vic Villager = villager{
		animal: animals.Bull.Name(),
		name:   vic}
)
