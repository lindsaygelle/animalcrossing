package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	apple string = "Apple"
)

var (
	// Apple is an Animal Crossing villager.
	//
	// Apple is a Hamster.
	Apple Villager = villager{
		animal: animals.Hamster.Name(),
		name:   apple}
)
