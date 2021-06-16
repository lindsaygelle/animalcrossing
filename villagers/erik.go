package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	erik string = "Erik"
)

var (
	// Erik is an Animal Crossing villager.
	//
	// Erik is a Deer.
	Erik Villager = villager{
		animal: animals.Deer.Name(),
		name:   erik}
)
