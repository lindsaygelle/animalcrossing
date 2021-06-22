package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	lief string = "Lief"
)

var (
	// Lief is an Animal Crossing villager.
	//
	// Lief is a Sloth.
	Lief Villager = villager{
		animal: animals.Sloth.Name(),
		name:   lief}
)
