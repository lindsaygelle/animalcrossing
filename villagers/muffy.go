package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	muffy string = "Muffy"
)

var (
	// Muffy is an Animal Crossing villager.
	//
	// Muffy is a Sheep.
	Muffy Villager = villager{
		animal: animals.Sheep.Name(),
		name:   muffy}
)
