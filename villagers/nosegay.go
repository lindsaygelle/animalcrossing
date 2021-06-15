package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	nosegay string = "Nosegay"
)

var (
	// Nosegay is an Animal Crossing villager.
	//
	// Nosegay is an Anteater.
	Nosegay Villager = villager{
		animal: animals.Anteater.Name(),
		name:   nosegay}
)
