package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	nobuo string = "Nobuo"
)

var (
	// Nobuo is an Animal Crossing villager.
	//
	// Nobuo is a Penguin.
	Nobuo Villager = villager{
		animal: animals.Penguin.Name(),
		name:   nobuo}
)
