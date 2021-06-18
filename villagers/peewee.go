package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	peewee string = "Peewee"
)

var (
	// Peewee is an Animal Crossing villager.
	//
	// Peewee is a Gorilla.
	Peewee Villager = villager{
		animal: animals.Gorilla.Name(),
		name:   peewee}
)
