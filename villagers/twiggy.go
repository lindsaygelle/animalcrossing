package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	twiggy string = "Twiggy"
)

var (
	// Twiggy is an Animal Crossing villager.
	//
	// Twiggy is a Bird.
	Twiggy Villager = villager{
		animal: animals.Bird.Name(),
		name:   twiggy}
)
