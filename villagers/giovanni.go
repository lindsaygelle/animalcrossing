package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	giovanni string = "Giovanni"
)

var (
	// Giovanni is an Animal Crossing villager.
	//
	// Giovanni is a Bird.
	Giovanni Villager = villager{
		animal: animals.Bird.Name(),
		name:   giovanni}
)
