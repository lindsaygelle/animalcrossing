package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	poko string = "Poko"
)

var (
	// Poko is an Animal Crossing animal.
	//
	// Poko is a Bear.
	Poko Villager = villager{
		animal: animals.Bear.Name(),
		name:   poko}
)
