package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	aurora string = "Aurora"
)

var (
	// Aurora is an Animal Crossing villager.
	//
	// Aurora is a Penguin.
	Aurora Villager = villager{
		animal: animals.Penguin.Name(),
		name:   aurora}
)
