package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	doc string = "Doc"
)

var (
	// Doc is an Animal Crossing villager.
	//
	// Doc is a Rabbit.
	Doc Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   doc}
)
