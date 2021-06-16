package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	goose string = "Goose"
)

var (
	// Goose is an Animal Crossing villager.
	//
	// Goose is a Chicken.
	Goose Villager = villager{
		animal: animals.Chicken.Name(),
		name:   goose}
)
