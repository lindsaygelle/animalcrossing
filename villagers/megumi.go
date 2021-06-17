package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	megumi string = "Megumi"
)

var (
	// Megumi is an Animal Crossing villager.
	//
	// Megumi is a Dog.
	Megumi Villager = villager{
		animal: animals.Dog.Name(),
		name:   megumi}
)
