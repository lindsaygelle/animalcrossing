package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cashmere string = "Cashmere"
)

var (
	// Cashmere is an Animal Crossing villager.
	//
	// Cashmere is a Sheep.
	Cashmere Villager = villager{
		animal: animals.Sheep.Name(),
		name:   cashmere}
)
