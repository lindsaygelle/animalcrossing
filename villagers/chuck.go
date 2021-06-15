package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	chuck string = "Chuck"
)

var (
	// Chuck is an Animal Crossing villager.
	//
	// Chuck is a Bull.
	Chuck Villager = villager{
		animal: animals.Bull.Name(),
		name:   chuck}
)
