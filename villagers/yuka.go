package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	yuka string = "Yuka"
)

var (
	// Yuka is an Animal Crossing villager.
	//
	// Yuka is a Koala.
	Yuka Villager = villager{
		animal: animals.Koala.Name(),
		name:   yuka}
)
