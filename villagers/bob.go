package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bob string = "Bob"
)

var (
	// Bob is an Animal Crossing villager.
	//
	// Bob is a Cat.
	Bob Villager = villager{
		animal: animals.Cat.Name(),
		name:   bob}
)
