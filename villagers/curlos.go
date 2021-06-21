package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	curlos string = "Curlos"
)

var (
	// Curlos is an Animal Crossing villager.
	//
	// Curlos is a Sheep.
	Curlos Villager = villager{
		animal: animals.Sheep.Name(),
		name:   curlos}
)
