package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	coco string = "Coco"
)

var (
	// Coco is an Animal Crossing villager.
	//
	// Coco is a Rabbit.
	Coco Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   coco}
)
