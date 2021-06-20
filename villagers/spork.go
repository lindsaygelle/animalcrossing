package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	spork string = "Spork"
)

var (
	// Spork is an Animal Crossing villager.
	//
	// Spork is a Pig.
	Spork Villager = villager{
		animal: animals.Pig.Name(),
		name:   spork}
)
