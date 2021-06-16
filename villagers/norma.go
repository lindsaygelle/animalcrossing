package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	norma string = "Norma"
)

var (
	// Norma is an Animal Crossing villager.
	//
	// Norma is a Cow.
	Norma Villager = villager{
		animal: animals.Cow.Name(),
		name:   norma}
)
