package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	vesta string = "Vesta"
)

var (
	// Vesta is an Animal Crossing villager.
	//
	// Vesta is a Sheep.
	Vesta Villager = villager{
		animal: animals.Sheep.Name(),
		name:   vesta}
)
