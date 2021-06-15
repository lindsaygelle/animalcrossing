package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kaitlin string = "Kaitlin"
)

var (
	// Kaitlin is an Animal Crossing villager.
	//
	// Kaitlin is a Cat.
	Kaitlin Villager = villager{
		animal: animals.Cat.Name(),
		name:   kaitlin}
)
