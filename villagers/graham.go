package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	graham string = "Graham"
)

var (
	// Graham is an Animal Crossing villager.
	//
	// Graham is a Hamster.
	Graham Villager = villager{
		animal: animals.Hamster.Name(),
		name:   graham}
)
