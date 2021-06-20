package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	leilani string = "Leilani"
)

var (
	// Leilani is an Animal Crossing villager.
	//
	// Leilani is a Kappa.
	Leilani Villager = villager{
		animal: animals.Kappa.Name(),
		name:   leilani}
)
