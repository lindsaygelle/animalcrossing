package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	willow string = "Willow"
)

var (
	// Willow is an Animal Crossing villager.
	//
	// Willow is a Sheep.
	Willow Villager = villager{
		animal: animals.Sheep.Name(),
		name:   willow}
)
