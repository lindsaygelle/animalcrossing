package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	raymond string = "Raymond"
)

var (
	// Raymond is an Animal Crossing villager.
	//
	// Raymond is a Cat.
	Raymond Villager = villager{
		animal: animals.Cat.Name(),
		name:   raymond}
)
