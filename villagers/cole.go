package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	cole string = "Cole"
)

var (
	// Cole is an Animal Crossing villager.
	//
	// Cole is a Rabbit.
	Cole Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   cole}
)
