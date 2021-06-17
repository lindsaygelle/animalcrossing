package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gloria string = "Gloria"
)

var (
	// Gloria is an Animal Crossing villager.
	//
	// Gloria is a Duck.
	Gloria Villager = villager{
		animal: animals.Duck.Name(),
		name:   gloria}
)
