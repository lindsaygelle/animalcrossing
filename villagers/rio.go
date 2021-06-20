package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	rio string = "Rio"
)

var (
	// Rio is an Animal Crossing villager.
	//
	// Rio is an Ostrich.
	Rio Villager = villager{
		animal: animals.Ostrich.Name(),
		name:   rio}
)
