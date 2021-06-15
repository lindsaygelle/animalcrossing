package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	murphy string = "Murphy"
)

var (
	// Murphy is an Animal Crossing animal.
	//
	// Murphy is a Bear.
	Murphy Villager = villager{
		animal: animals.Bear.Name(),
		name:   murphy}
)
