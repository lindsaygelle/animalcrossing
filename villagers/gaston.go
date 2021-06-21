package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gaston string = "Gaston"
)

var (
	// Gaston is an Animal Crossing villager.
	//
	// Gaston is a Rabbit.
	Gaston Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   gaston}
)
