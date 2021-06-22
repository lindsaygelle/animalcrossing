package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kicks string = "Kicks"
)

var (
	Kicks Villager = villager{
		animal: animals.Skunk.Name(),
		name:   kicks}
)
