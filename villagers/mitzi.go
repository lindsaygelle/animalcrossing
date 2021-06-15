package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	mitzi string = "Mitzi"
)

var (
	// Mitzi is an Animal Crossing villager.
	//
	// Mitzi is a Cat.
	Mitzi Villager = villager{
		animal: animals.Cat.Name(),
		name:   mitzi}
)
