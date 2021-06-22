package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	caroline string = "Caroline"
)

var (
	// Caroline is an Animal Crossing villager.
	//
	// Caroline is a Squirrel.
	Caroline Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   caroline}
)
