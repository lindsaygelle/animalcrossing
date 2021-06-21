package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	snake string = "Snake"
)

var (
	// Snake is an Animal Crossing villager.
	//
	// Snake is a Rabbit.
	Snake Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   snake}
)
