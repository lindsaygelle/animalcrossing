package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	flash string = "Flash"
)

var (
	// Flash is an Animal Crossing villager.
	//
	// Flash is a Bird.
	Flash Villager = villager{
		animal: animals.Bird.Name(),
		name:   flash}
)
