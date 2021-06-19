package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	epona string = "Epona"
)

var (
	// Epona is an Animal Crossing villager.
	//
	// Epona is a Horse.
	Epona Villager = villager{
		animal: animals.Horse.Name(),
		name:   epona}
)
