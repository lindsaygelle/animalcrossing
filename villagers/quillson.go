package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	quillson string = "Quillson"
)

var (
	// Quillson is an Animal Crossing villager.
	//
	// Quillson is a Duck.
	Quillson Villager = villager{
		animal: animals.Duck.Name(),
		name:   quillson}
)
