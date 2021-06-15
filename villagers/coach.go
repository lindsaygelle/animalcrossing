package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	coach string = "Coach"
)

var (
	// Coach is an Animal Crossing villager.
	//
	// Coach is a Bull.
	Coach Villager = villager{
		animal: animals.Bull.Name(),
		name:   coach}
)
