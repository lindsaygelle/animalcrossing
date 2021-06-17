package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	shinabiru string = "Shinabiru"
)

var (
	// Shinabiru is an Animal Crossing villager.
	//
	// Shinabiru is a Duck.
	Shinabiru Villager = villager{
		animal: animals.Duck.Name(),
		name:   shinabiru}
)
