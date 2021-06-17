package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	jambette string = "Jambette"
)

var (
	// Jambette is an Animal Crossing villager.
	//
	// Jambette is a Frog.
	Jambette Villager = villager{
		animal: animals.Frog.Name(),
		name:   jambette}
)
