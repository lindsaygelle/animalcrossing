package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	flurry string = "Flurry"
)

var (
	// Flurry is an Animal Crossing villager.
	//
	// Flurry is a Hamster.
	Flurry Villager = villager{
		animal: animals.Hamster.Name(),
		name:   flurry}
)
