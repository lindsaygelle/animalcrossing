package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	dom string = "Dom"
)

var (
	// Dom is an Animal Crossing villager.
	//
	// Dom is a Sheep.
	Dom Villager = villager{
		animal: animals.Sheep.Name(),
		name:   dom}
)
