package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	becky string = "Becky"
)

var (
	// Becky is an Animal Crossing villager.
	//
	// Becky is a Chicken.
	Becky Villager = villager{
		animal: animals.Chicken.Name(),
		name:   becky}
)
