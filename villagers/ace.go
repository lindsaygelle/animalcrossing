package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ace string = "Ace"
)

var (
	// Ace is an Animal Crossing villager.
	//
	// Ace is a Bird.
	Ace Villager = villager{
		animal: animals.Bird.Name(),
		name:   ace}
)
