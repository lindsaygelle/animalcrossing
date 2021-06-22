package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	nibbles string = "Nibbles"
)

var (
	// Nibbles is an Animal Crossing villager.
	//
	// Nibbles is a Squirrel.
	Nibbles Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   nibbles}
)
