package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	katt string = "Katt"
)

var (
	// Katt is an Animal Crossing villager.
	//
	// Katt is a Cat.
	Katt Villager = villager{
		animal: animals.Cat.Name(),
		name:   katt}
)
