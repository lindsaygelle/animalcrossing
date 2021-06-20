package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	blathers string = "Blathers"
)

var (
	// Blathers is an Animal Crossing villager.
	//
	// Blathers is an Owl.
	Blathers Villager = villager{
		animal: animals.Owl.Name(),
		name:   blathers}
)
